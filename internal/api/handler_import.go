package api

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/repository"
)

// ImportExportHandler handles import/export API endpoints.
type ImportExportHandler struct {
	addressRepo *repository.AddressRepository
	subnetRepo  *repository.SubnetRepository
}

// NewImportExportHandler creates a new ImportExportHandler.
func NewImportExportHandler(addressRepo *repository.AddressRepository, subnetRepo *repository.SubnetRepository) *ImportExportHandler {
	return &ImportExportHandler{
		addressRepo: addressRepo,
		subnetRepo:  subnetRepo,
	}
}

// Import handles POST /api/v1/import/{type}
func (h *ImportExportHandler) Import(w http.ResponseWriter, r *http.Request) {
	importType := chi.URLParam(r, "type")
	if importType == "" {
		BadRequest(w, r, errcode.ValidationError, "import type is required")
		return
	}

	// Limit upload size to 10MB
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)

	reader := csv.NewReader(r.Body)
	reader.TrimLeadingSpace = true

	// Read header row
	headers, err := reader.Read()
	if err != nil {
		BadRequest(w, r, errcode.ImportValidationFailed, "failed to read CSV header")
		return
	}

	var imported, skipped, errors int

	switch importType {
	case "addresses":
		imported, skipped, errors, err = h.importAddresses(r, reader, headers)
	case "subnets":
		imported, skipped, errors, err = h.importSubnets(r, reader, headers)
	default:
		BadRequest(w, r, errcode.ValidationError, fmt.Sprintf("unsupported import type: %s", importType))
		return
	}

	if err != nil {
		InternalError(w, r, "import failed")
		return
	}

	OK(w, r, map[string]interface{}{
		"imported": imported,
		"skipped":  skipped,
		"errors":   errors,
	})
}

// Export handles GET /api/v1/export/{type}
func (h *ImportExportHandler) Export(w http.ResponseWriter, r *http.Request) {
	exportType := chi.URLParam(r, "type")
	if exportType == "" {
		BadRequest(w, r, errcode.ValidationError, "export type is required")
		return
	}

	switch exportType {
	case "addresses":
		h.exportAddresses(w, r)
	case "subnets":
		h.exportSubnets(w, r)
	default:
		BadRequest(w, r, errcode.ValidationError, fmt.Sprintf("unsupported export type: %s", exportType))
		return
	}
}

func (h *ImportExportHandler) importAddresses(r *http.Request, reader *csv.Reader, headers []string) (imported, skipped, errors int, err error) {
	// Build column index map
	colMap := make(map[string]int)
	for i, h := range headers {
		colMap[h] = i
	}

	// Required: subnet_id, ip
	if _, ok := colMap["subnet_id"]; !ok {
		return 0, 0, 1, fmt.Errorf("missing required column: subnet_id")
	}
	if _, ok := colMap["ip"]; !ok {
		return 0, 0, 1, fmt.Errorf("missing required column: ip")
	}

	for {
		row, readErr := reader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			errors++
			continue
		}

		getCol := func(name string) string {
			if idx, ok := colMap[name]; ok && idx < len(row) {
				return row[idx]
			}
			return ""
		}

		subnetID, parseErr := strconv.ParseInt(getCol("subnet_id"), 10, 64)
		if parseErr != nil {
			slog.Warn("import address: invalid subnet_id", "error", parseErr)
			skipped++
			continue
		}

		ipText := getCol("ip")
		if ipText == "" {
			skipped++
			continue
		}

		// Check if already exists
		existing, _ := h.addressRepo.GetByIPText(r.Context(), subnetID, ipText)
		if existing != nil {
			slog.Warn("import address: address already exists", "ip", ipText, "subnet_id", subnetID)
			skipped++
			continue
		}

		ipBytes, parseErr := ipam.IPToBytes(ipText)
		if parseErr != nil {
			slog.Warn("import address: invalid IP", "ip", ipText, "error", parseErr)
			skipped++
			continue
		}

		hostname := getCol("hostname")
		mac := getCol("mac")
		description := getCol("description")
		status := getCol("status")
		if status == "" {
			status = "active"
		}
		owner := getCol("owner")

		_, createErr := h.addressRepo.Create(
			r.Context(), subnetID, ipBytes, ipText,
			hostname, mac, description, status,
			nil, false, nil, "", owner,
		)
		if createErr != nil {
			errors++
			continue
		}
		imported++
	}

	return imported, skipped, errors, nil
}

func (h *ImportExportHandler) importSubnets(r *http.Request, reader *csv.Reader, headers []string) (imported, skipped, errors int, err error) {
	colMap := make(map[string]int)
	for i, h := range headers {
		colMap[h] = i
	}

	if _, ok := colMap["section_id"]; !ok {
		return 0, 0, 1, fmt.Errorf("missing required column: section_id")
	}
	if _, ok := colMap["cidr"]; !ok {
		return 0, 0, 1, fmt.Errorf("missing required column: cidr")
	}

	for {
		row, readErr := reader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			errors++
			continue
		}

		getCol := func(name string) string {
			if idx, ok := colMap[name]; ok && idx < len(row) {
				return row[idx]
			}
			return ""
		}

		sectionID, parseErr := strconv.ParseInt(getCol("section_id"), 10, 64)
		if parseErr != nil {
			slog.Warn("import subnet: invalid section_id", "error", parseErr)
			skipped++
			continue
		}

		cidr := getCol("cidr")
		if cidr == "" {
			skipped++
			continue
		}

		if validateErr := ipam.ValidateCIDR(cidr); validateErr != nil {
			slog.Warn("import subnet: invalid CIDR", "cidr", cidr, "error", validateErr)
			skipped++
			continue
		}

		info, parseErr := ipam.ParseCIDR(cidr)
		if parseErr != nil {
			skipped++
			continue
		}

		// Extract pure IP from CIDR (info.Network is "192.168.1.0/24")
		_, ipNet, _ := net.ParseCIDR(info.Network)
		networkBytes := ipNet.IP
		if v4 := networkBytes.To4(); v4 != nil {
			networkBytes = v4
		}

		name := getCol("name")
		description := getCol("description")
		status := getCol("status")
		if status == "" {
			status = "active"
		}

		_, createErr := h.subnetRepo.Create(
			r.Context(), sectionID, nil, info.Network,
			uint8(info.IPVersion), networkBytes, uint8(info.PrefixLen),
			name, description, nil, nil, nil, nil,
			false, false, 0, false, status,
		)
		if createErr != nil {
			errors++
			continue
		}
		imported++
	}

	return imported, skipped, errors, nil
}

func (h *ImportExportHandler) exportAddresses(w http.ResponseWriter, r *http.Request) {
	subnetIDStr := r.URL.Query().Get("subnet_id")
	if subnetIDStr == "" {
		BadRequest(w, r, errcode.ValidationError, "subnet_id query parameter is required")
		return
	}

	subnetID, err := strconv.ParseInt(subnetIDStr, 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet_id")
		return
	}

	addresses, err := h.addressRepo.List(r.Context(), subnetID, 10000, 0)
	if err != nil {
		InternalError(w, r, "failed to list addresses")
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=addresses.csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"subnet_id", "ip", "hostname", "mac", "description", "status", "owner"})

	// Write rows
	for _, addr := range addresses {
		writer.Write([]string{
			strconv.FormatInt(addr.SubnetID, 10),
			addr.IpText,
			addr.Hostname,
			addr.Mac,
			addr.Description,
			addr.Status,
			addr.Owner,
		})
	}
}

func (h *ImportExportHandler) exportSubnets(w http.ResponseWriter, r *http.Request) {
	sectionIDStr := r.URL.Query().Get("section_id")
	if sectionIDStr == "" {
		BadRequest(w, r, errcode.ValidationError, "section_id query parameter is required")
		return
	}

	sectionID, err := strconv.ParseInt(sectionIDStr, 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section_id")
		return
	}

	subnets, err := h.subnetRepo.List(r.Context(), sectionID, 10000, 0)
	if err != nil {
		InternalError(w, r, "failed to list subnets")
		return
	}

	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=subnets.csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"section_id", "cidr", "name", "description", "status"})

	// Write rows
	for _, subnet := range subnets {
		writer.Write([]string{
			strconv.FormatInt(subnet.SectionID, 10),
			subnet.Cidr,
			subnet.Name,
			subnet.Description,
			subnet.Status,
		})
	}
}
