package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/jincaiw/goipam/internal/api"
	"github.com/jincaiw/goipam/internal/config"
	"github.com/jincaiw/goipam/internal/rbac"
)

// Server represents the HTTP server.
type Server struct {
	cfg                 *config.Config
	mux                 *chi.Mux
	authHandler         *api.AuthHandler
	userHandler         *api.UserHandler
	sectionHandler      *api.SectionHandler
	folderHandler       *api.FolderHandler
	subnetHandler       *api.SubnetHandler
	addressHandler      *api.AddressHandler
	vlanHandler         *api.VlanHandler
	vrfHandler          *api.VrfHandler
	deviceHandler       *api.DeviceHandler
	locationHandler     *api.LocationHandler
	natHandler          *api.NatHandler
	nameserverHandler   *api.NameserverHandler
	scanHandler         *api.ScanHandler
	requestHandler      *api.RequestHandler
	searchHandler       *api.SearchHandler
	auditHandler        *api.AuditHandler
	importExportHandler *api.ImportExportHandler
	dashboardHandler    *api.DashboardHandler
	roleHandler         *api.RoleHandler
	groupHandler        *api.GroupHandler
	settingsHandler     *api.SettingsHandler
	rbacMW              *rbac.Middleware
}

// Option is a functional option for Server.
type Option func(*Server)

// WithAuthHandler sets the auth handler.
func WithAuthHandler(h *api.AuthHandler) Option {
	return func(s *Server) { s.authHandler = h }
}

// WithUserHandler sets the user handler.
func WithUserHandler(h *api.UserHandler) Option {
	return func(s *Server) { s.userHandler = h }
}

// WithSectionHandler sets the section handler.
func WithSectionHandler(h *api.SectionHandler) Option {
	return func(s *Server) { s.sectionHandler = h }
}

// WithFolderHandler sets the folder handler.
func WithFolderHandler(h *api.FolderHandler) Option {
	return func(s *Server) { s.folderHandler = h }
}

// WithSubnetHandler sets the subnet handler.
func WithSubnetHandler(h *api.SubnetHandler) Option {
	return func(s *Server) { s.subnetHandler = h }
}

// WithAddressHandler sets the address handler.
func WithAddressHandler(h *api.AddressHandler) Option {
	return func(s *Server) { s.addressHandler = h }
}

// WithNameserverHandler sets the nameserver handler.
func WithNameserverHandler(h *api.NameserverHandler) Option {
	return func(s *Server) { s.nameserverHandler = h }
}

// WithScanHandler sets the scan handler.
func WithScanHandler(h *api.ScanHandler) Option {
	return func(s *Server) { s.scanHandler = h }
}

// WithRequestHandler sets the request handler.
func WithRequestHandler(h *api.RequestHandler) Option {
	return func(s *Server) { s.requestHandler = h }
}

// WithSearchHandler sets the search handler.
func WithSearchHandler(h *api.SearchHandler) Option {
	return func(s *Server) { s.searchHandler = h }
}

// WithAuditHandler sets the audit handler.
func WithAuditHandler(h *api.AuditHandler) Option {
	return func(s *Server) { s.auditHandler = h }
}

// WithImportExportHandler sets the import/export handler.
func WithImportExportHandler(h *api.ImportExportHandler) Option {
	return func(s *Server) { s.importExportHandler = h }
}

// WithVlanHandler sets the VLAN handler.
func WithVlanHandler(h *api.VlanHandler) Option {
	return func(s *Server) { s.vlanHandler = h }
}

// WithVrfHandler sets the VRF handler.
func WithVrfHandler(h *api.VrfHandler) Option {
	return func(s *Server) { s.vrfHandler = h }
}

// WithDeviceHandler sets the device handler.
func WithDeviceHandler(h *api.DeviceHandler) Option {
	return func(s *Server) { s.deviceHandler = h }
}

// WithLocationHandler sets the location handler.
func WithLocationHandler(h *api.LocationHandler) Option {
	return func(s *Server) { s.locationHandler = h }
}

// WithNatHandler sets the NAT handler.
func WithNatHandler(h *api.NatHandler) Option {
	return func(s *Server) { s.natHandler = h }
}

// WithDashboardHandler sets the dashboard handler.
func WithDashboardHandler(h *api.DashboardHandler) Option {
	return func(s *Server) { s.dashboardHandler = h }
}

// WithRoleHandler sets the role handler.
func WithRoleHandler(h *api.RoleHandler) Option {
	return func(s *Server) { s.roleHandler = h }
}

// WithGroupHandler sets the group handler.
func WithGroupHandler(h *api.GroupHandler) Option {
	return func(s *Server) { s.groupHandler = h }
}

// WithSettingsHandler sets the settings handler.
func WithSettingsHandler(h *api.SettingsHandler) Option {
	return func(s *Server) { s.settingsHandler = h }
}

// WithRBACMiddleware sets the RBAC middleware.
func WithRBACMiddleware(mw *rbac.Middleware) Option {
	return func(s *Server) { s.rbacMW = mw }
}

// New creates a new Server with routes and middleware configured.
func New(cfg *config.Config, opts ...Option) *Server {
	s := &Server{
		cfg: cfg,
		mux: chi.NewRouter(),
	}

	for _, opt := range opts {
		opt(s)
	}

	s.setupMiddleware()
	s.setupRoutes()
	s.setupStaticFiles()

	return s
}

func (s *Server) setupMiddleware() {
	s.mux.Use(chimw.RequestID)
	s.mux.Use(chimw.RealIP)
	s.mux.Use(requestIDContext)
	s.mux.Use(chimw.Logger)
	s.mux.Use(chimw.Recoverer)
	s.mux.Use(chimw.Timeout(60 * time.Second))
	s.mux.Use(chimw.Compress(5))
	s.mux.Use(corsMiddleware)
}

func (s *Server) setupRoutes() {
	// Health & metrics (no auth required)
	s.mux.Get("/healthz", s.healthz)
	s.mux.Get("/readyz", s.readyz)
	s.setupMetrics(s.mux)

	// API v1 routes
	s.mux.Route("/api/v1", func(r chi.Router) {
		// Auth routes (public)
		if s.authHandler != nil {
			r.Post("/auth/login", s.authHandler.Login)
			r.Post("/auth/refresh", s.authHandler.RefreshToken)
		} else {
			r.Post("/auth/login", s.placeholderHandler("auth.login"))
			r.Post("/auth/refresh", s.placeholderHandler("auth.refresh"))
		}

		// Authenticated routes
		r.Group(func(r chi.Router) {
			if s.rbacMW != nil {
				r.Use(s.rbacMW.RequireAuth)
			}

			if s.authHandler != nil {
				r.Post("/auth/logout", s.authHandler.Logout)
				r.Get("/auth/me", s.authHandler.Me)
				r.Post("/auth/change-password", s.authHandler.ChangePassword)
			} else {
				r.Post("/auth/logout", s.placeholderHandler("auth.logout"))
				r.Get("/auth/me", s.placeholderHandler("auth.me"))
				r.Post("/auth/change-password", s.placeholderHandler("auth.change-password"))
			}

			// Sections
			if s.sectionHandler != nil {
				r.Get("/sections", s.sectionHandler.List)
				r.Post("/sections", s.sectionHandler.Create)
				r.Get("/sections/{id}", s.sectionHandler.Get)
				r.Put("/sections/{id}", s.sectionHandler.Update)
				r.Delete("/sections/{id}", s.sectionHandler.Delete)
				r.Get("/sections/{id}/tree", s.sectionHandler.Tree)
				r.Get("/sections/{id}/stats", s.sectionHandler.Stats)
			} else {
				r.Get("/sections", s.placeholderHandler("sections.list"))
				r.Post("/sections", s.placeholderHandler("sections.create"))
				r.Get("/sections/{id}", s.placeholderHandler("sections.get"))
				r.Put("/sections/{id}", s.placeholderHandler("sections.update"))
				r.Delete("/sections/{id}", s.placeholderHandler("sections.delete"))
				r.Get("/sections/{id}/tree", s.placeholderHandler("sections.tree"))
				r.Get("/sections/{id}/stats", s.placeholderHandler("sections.stats"))
			}

			// Folders
			if s.folderHandler != nil {
				r.Get("/folders", s.folderHandler.List)
				r.Post("/folders", s.folderHandler.Create)
				r.Get("/folders/{id}", s.folderHandler.Get)
				r.Put("/folders/{id}", s.folderHandler.Update)
				r.Delete("/folders/{id}", s.folderHandler.Delete)
			} else {
				r.Get("/folders", s.placeholderHandler("folders.list"))
				r.Post("/folders", s.placeholderHandler("folders.create"))
				r.Get("/folders/{id}", s.placeholderHandler("folders.get"))
				r.Put("/folders/{id}", s.placeholderHandler("folders.update"))
				r.Delete("/folders/{id}", s.placeholderHandler("folders.delete"))
			}

			// Subnets
			if s.subnetHandler != nil {
				r.Get("/subnets", s.subnetHandler.List)
				r.Post("/subnets", s.subnetHandler.Create)
				r.Get("/subnets/{id}", s.subnetHandler.Get)
				r.Put("/subnets/{id}", s.subnetHandler.Update)
				r.Delete("/subnets/{id}", s.subnetHandler.Delete)
				r.Get("/subnets/{id}/children", s.subnetHandler.Children)
				r.Get("/subnets/{id}/addresses", s.subnetHandler.Addresses)
				r.Get("/subnets/{id}/first-free-ip", s.subnetHandler.FirstFreeIP)
				r.Post("/subnets/{id}/split", s.subnetHandler.Split)
				r.Get("/subnets/{id}/available-ranges", s.subnetHandler.AvailableRanges)
				r.Get("/subnets/{id}/stats", s.subnetHandler.Stats)
				r.Post("/subnets/check-overlap", s.subnetHandler.CheckOverlap)
			} else {
				r.Get("/subnets", s.placeholderHandler("subnets.list"))
				r.Post("/subnets", s.placeholderHandler("subnets.create"))
				r.Get("/subnets/{id}", s.placeholderHandler("subnets.get"))
				r.Put("/subnets/{id}", s.placeholderHandler("subnets.update"))
				r.Delete("/subnets/{id}", s.placeholderHandler("subnets.delete"))
				r.Get("/subnets/{id}/children", s.placeholderHandler("subnets.children"))
				r.Get("/subnets/{id}/addresses", s.placeholderHandler("subnets.addresses"))
				r.Get("/subnets/{id}/first-free-ip", s.placeholderHandler("subnets.firstFreeIP"))
				r.Post("/subnets/{id}/split", s.placeholderHandler("subnets.split"))
				r.Get("/subnets/{id}/available-ranges", s.placeholderHandler("subnets.availableRanges"))
				r.Get("/subnets/{id}/stats", s.placeholderHandler("subnets.stats"))
				r.Post("/subnets/check-overlap", s.placeholderHandler("subnets.checkOverlap"))
			}

			// Addresses
			if s.addressHandler != nil {
				r.Get("/addresses", s.addressHandler.List)
				r.Post("/addresses", s.addressHandler.Create)
				r.Post("/addresses/first-free", s.addressHandler.FirstFree)
				r.Get("/addresses/{id}", s.addressHandler.Get)
				r.Put("/addresses/{id}", s.addressHandler.Update)
				r.Delete("/addresses/{id}", s.addressHandler.Delete)
				r.Post("/addresses/{id}/ping", s.addressHandler.Ping)
				r.Post("/addresses/{id}/dns-resolve", s.addressHandler.DnsResolve)
			} else {
				r.Get("/addresses", s.placeholderHandler("addresses.list"))
				r.Post("/addresses", s.placeholderHandler("addresses.create"))
				r.Post("/addresses/first-free", s.placeholderHandler("addresses.firstFree"))
				r.Get("/addresses/{id}", s.placeholderHandler("addresses.get"))
				r.Put("/addresses/{id}", s.placeholderHandler("addresses.update"))
				r.Delete("/addresses/{id}", s.placeholderHandler("addresses.delete"))
				r.Post("/addresses/{id}/ping", s.placeholderHandler("addresses.ping"))
				r.Post("/addresses/{id}/dns-resolve", s.placeholderHandler("addresses.dnsResolve"))
			}

			// VLANs & L2 Domains
			if s.vlanHandler != nil {
				r.Get("/vlans", s.vlanHandler.ListVlans)
				r.Post("/vlans", s.vlanHandler.CreateVlan)
				r.Get("/vlans/{id}", s.vlanHandler.GetVlan)
				r.Put("/vlans/{id}", s.vlanHandler.UpdateVlan)
				r.Delete("/vlans/{id}", s.vlanHandler.DeleteVlan)
				r.Get("/l2-domains", s.vlanHandler.ListL2Domains)
				r.Post("/l2-domains", s.vlanHandler.CreateL2Domain)
			} else {
				r.Get("/vlans", s.placeholderHandler("vlans.list"))
				r.Post("/vlans", s.placeholderHandler("vlans.create"))
				r.Get("/vlans/{id}", s.placeholderHandler("vlans.get"))
				r.Put("/vlans/{id}", s.placeholderHandler("vlans.update"))
				r.Delete("/vlans/{id}", s.placeholderHandler("vlans.delete"))
				r.Get("/l2-domains", s.placeholderHandler("l2domains.list"))
				r.Post("/l2-domains", s.placeholderHandler("l2domains.create"))
			}

			// VRFs
			if s.vrfHandler != nil {
				r.Get("/vrfs", s.vrfHandler.List)
				r.Post("/vrfs", s.vrfHandler.Create)
				r.Get("/vrfs/{id}", s.vrfHandler.Get)
				r.Put("/vrfs/{id}", s.vrfHandler.Update)
				r.Delete("/vrfs/{id}", s.vrfHandler.Delete)
			} else {
				r.Get("/vrfs", s.placeholderHandler("vrfs.list"))
				r.Post("/vrfs", s.placeholderHandler("vrfs.create"))
				r.Get("/vrfs/{id}", s.placeholderHandler("vrfs.get"))
				r.Put("/vrfs/{id}", s.placeholderHandler("vrfs.update"))
				r.Delete("/vrfs/{id}", s.placeholderHandler("vrfs.delete"))
			}

			// Devices
			if s.deviceHandler != nil {
				r.Get("/devices", s.deviceHandler.ListDevices)
				r.Post("/devices", s.deviceHandler.CreateDevice)
				r.Get("/devices/{id}", s.deviceHandler.GetDevice)
				r.Put("/devices/{id}", s.deviceHandler.UpdateDevice)
				r.Delete("/devices/{id}", s.deviceHandler.DeleteDevice)
			} else {
				r.Get("/devices", s.placeholderHandler("devices.list"))
				r.Post("/devices", s.placeholderHandler("devices.create"))
				r.Get("/devices/{id}", s.placeholderHandler("devices.get"))
				r.Put("/devices/{id}", s.placeholderHandler("devices.update"))
				r.Delete("/devices/{id}", s.placeholderHandler("devices.delete"))
			}

			// Device Types
			if s.deviceHandler != nil {
				r.Get("/device-types", s.deviceHandler.ListDeviceTypes)
				r.Post("/device-types", s.deviceHandler.CreateDeviceType)
			} else {
				r.Get("/device-types", s.placeholderHandler("deviceTypes.list"))
				r.Post("/device-types", s.placeholderHandler("deviceTypes.create"))
			}

			// Locations
			if s.locationHandler != nil {
				r.Get("/locations", s.locationHandler.ListLocations)
				r.Post("/locations", s.locationHandler.CreateLocation)
				r.Get("/locations/{id}", s.locationHandler.GetLocation)
				r.Put("/locations/{id}", s.locationHandler.UpdateLocation)
				r.Delete("/locations/{id}", s.locationHandler.DeleteLocation)
			} else {
				r.Get("/locations", s.placeholderHandler("locations.list"))
				r.Post("/locations", s.placeholderHandler("locations.create"))
				r.Get("/locations/{id}", s.placeholderHandler("locations.get"))
				r.Put("/locations/{id}", s.placeholderHandler("locations.update"))
				r.Delete("/locations/{id}", s.placeholderHandler("locations.delete"))
			}

			// Racks
			if s.locationHandler != nil {
				r.Get("/racks", s.locationHandler.ListRacks)
				r.Post("/racks", s.locationHandler.CreateRack)
				r.Get("/racks/{id}", s.locationHandler.GetRack)
				r.Put("/racks/{id}", s.locationHandler.UpdateRack)
				r.Delete("/racks/{id}", s.locationHandler.DeleteRack)
				r.Post("/racks/{id}/mount-device", s.locationHandler.MountDevice)
				r.Post("/racks/{id}/unmount-device", s.locationHandler.UnmountDevice)
			} else {
				r.Get("/racks", s.placeholderHandler("racks.list"))
				r.Post("/racks", s.placeholderHandler("racks.create"))
				r.Get("/racks/{id}", s.placeholderHandler("racks.get"))
				r.Put("/racks/{id}", s.placeholderHandler("racks.update"))
				r.Delete("/racks/{id}", s.placeholderHandler("racks.delete"))
				r.Post("/racks/{id}/mount-device", s.placeholderHandler("racks.mount"))
				r.Post("/racks/{id}/unmount-device", s.placeholderHandler("racks.unmount"))
			}

			// NATs
			if s.natHandler != nil {
				r.Get("/nats", s.natHandler.List)
				r.Post("/nats", s.natHandler.Create)
				r.Get("/nats/{id}", s.natHandler.Get)
				r.Put("/nats/{id}", s.natHandler.Update)
				r.Delete("/nats/{id}", s.natHandler.Delete)
			} else {
				r.Get("/nats", s.placeholderHandler("nats.list"))
				r.Post("/nats", s.placeholderHandler("nats.create"))
				r.Get("/nats/{id}", s.placeholderHandler("nats.get"))
				r.Put("/nats/{id}", s.placeholderHandler("nats.update"))
				r.Delete("/nats/{id}", s.placeholderHandler("nats.delete"))
			}

			// Nameservers
			if s.nameserverHandler != nil {
				r.Get("/nameservers", s.nameserverHandler.List)
				r.Post("/nameservers", s.nameserverHandler.Create)
				r.Get("/nameservers/{id}", s.nameserverHandler.Get)
				r.Put("/nameservers/{id}", s.nameserverHandler.Update)
				r.Delete("/nameservers/{id}", s.nameserverHandler.Delete)
			} else {
				r.Get("/nameservers", s.placeholderHandler("nameservers.list"))
				r.Post("/nameservers", s.placeholderHandler("nameservers.create"))
				r.Get("/nameservers/{id}", s.placeholderHandler("nameservers.get"))
				r.Put("/nameservers/{id}", s.placeholderHandler("nameservers.update"))
				r.Delete("/nameservers/{id}", s.placeholderHandler("nameservers.delete"))
			}

			// Scans
			if s.scanHandler != nil {
				r.Post("/scans/subnets/{id}", s.scanHandler.CreateScan)
				r.Get("/scans/tasks", s.scanHandler.ListTasks)
				r.Get("/scans/tasks/{id}", s.scanHandler.GetTask)
				r.Post("/scans/tasks/{id}/cancel", s.scanHandler.CancelTask)
			} else {
				r.Post("/scans/subnets/{id}", s.placeholderHandler("scans.create"))
				r.Get("/scans/tasks", s.placeholderHandler("scans.list"))
				r.Get("/scans/tasks/{id}", s.placeholderHandler("scans.get"))
				r.Post("/scans/tasks/{id}/cancel", s.placeholderHandler("scans.cancel"))
			}

			// IP Requests
			if s.requestHandler != nil {
				r.Get("/requests", s.requestHandler.List)
				r.Post("/requests", s.requestHandler.Create)
				r.Post("/requests/{id}/approve", s.requestHandler.Approve)
				r.Post("/requests/{id}/reject", s.requestHandler.Reject)
			} else {
				r.Get("/requests", s.placeholderHandler("requests.list"))
				r.Post("/requests", s.placeholderHandler("requests.create"))
				r.Post("/requests/{id}/approve", s.placeholderHandler("requests.approve"))
				r.Post("/requests/{id}/reject", s.placeholderHandler("requests.reject"))
			}

			// Users
			if s.userHandler != nil {
				r.Get("/users", s.userHandler.List)
				r.Post("/users", s.userHandler.Create)
				r.Get("/users/{id}", s.userHandler.Get)
				r.Put("/users/{id}", s.userHandler.Update)
				r.Delete("/users/{id}", s.userHandler.Delete)
			} else {
				r.Get("/users", s.placeholderHandler("users.list"))
				r.Post("/users", s.placeholderHandler("users.create"))
				r.Get("/users/{id}", s.placeholderHandler("users.get"))
				r.Put("/users/{id}", s.placeholderHandler("users.update"))
				r.Delete("/users/{id}", s.placeholderHandler("users.delete"))
			}

			// Roles
			if s.roleHandler != nil {
				r.Get("/roles", s.roleHandler.List)
				r.Post("/roles", s.roleHandler.Create)
				r.Get("/roles/{id}", s.roleHandler.Get)
				r.Put("/roles/{id}", s.roleHandler.Update)
				r.Delete("/roles/{id}", s.roleHandler.Delete)
			} else {
				r.Get("/roles", s.placeholderHandler("roles.list"))
				r.Post("/roles", s.placeholderHandler("roles.create"))
				r.Get("/roles/{id}", s.placeholderHandler("roles.get"))
				r.Put("/roles/{id}", s.placeholderHandler("roles.update"))
				r.Delete("/roles/{id}", s.placeholderHandler("roles.delete"))
			}

			// Groups
			if s.groupHandler != nil {
				r.Get("/groups", s.groupHandler.List)
				r.Post("/groups", s.groupHandler.Create)
				r.Get("/groups/{id}", s.groupHandler.Get)
				r.Put("/groups/{id}", s.groupHandler.Update)
				r.Delete("/groups/{id}", s.groupHandler.Delete)
			} else {
				r.Get("/groups", s.placeholderHandler("groups.list"))
				r.Post("/groups", s.placeholderHandler("groups.create"))
				r.Get("/groups/{id}", s.placeholderHandler("groups.get"))
				r.Put("/groups/{id}", s.placeholderHandler("groups.update"))
				r.Delete("/groups/{id}", s.placeholderHandler("groups.delete"))
			}

			// Audit logs
			if s.auditHandler != nil {
				r.Get("/audit-logs", s.auditHandler.List)
			} else {
				r.Get("/audit-logs", s.placeholderHandler("audit.list"))
			}

			// System
			r.Get("/system/health", s.placeholderHandler("system.health"))

			// Import/Export
			if s.importExportHandler != nil {
				r.Post("/import/{type}", s.importExportHandler.Import)
				r.Get("/export/{type}", s.importExportHandler.Export)
			} else {
				r.Post("/import/{type}", s.placeholderHandler("import"))
				r.Get("/export/{type}", s.placeholderHandler("export"))
			}

			// Dashboard
			if s.dashboardHandler != nil {
				r.Get("/dashboard", s.dashboardHandler.Dashboard)
			} else {
				r.Get("/dashboard", s.placeholderHandler("dashboard"))
			}

			// Search
			if s.searchHandler != nil {
				r.Get("/search", s.searchHandler.Search)
			} else {
				r.Get("/search", s.placeholderHandler("search"))
			}

			// Settings
			if s.settingsHandler != nil {
				r.Get("/settings/{category}", s.settingsHandler.Get)
				r.Put("/settings/{category}", s.settingsHandler.Save)
			} else {
				r.Get("/settings/{category}", s.placeholderHandler("settings.get"))
				r.Put("/settings/{category}", s.placeholderHandler("settings.save"))
			}
		})
	})
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.cfg.Server.Listen, s.mux)
}

// Handler returns the underlying chi.Mux for testing.
func (s *Server) Handler() http.Handler {
	return s.mux
}
