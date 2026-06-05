package ipam

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
)

// IPVersion represents IP protocol version.
type IPVersion uint8

const (
	IPv4 IPVersion = 4
	IPv6 IPVersion = 6
)

// IPRange represents a range of IP addresses.
type IPRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Count int64  `json:"count"`
}

// SubnetInfo contains parsed subnet information.
type SubnetInfo struct {
	Network         string    `json:"network"`
	PrefixLen       int       `json:"prefix_len"`
	IPVersion       IPVersion `json:"ip_version"`
	FirstIP         string    `json:"first_ip"`
	LastIP          string    `json:"last_ip"`
	TotalAddresses  int64     `json:"total_addresses"`
	UsableAddresses int64     `json:"usable_addresses"`
	Mask            string    `json:"mask"`
}

// ParseCIDR parses a CIDR notation string and returns SubnetInfo.
func ParseCIDR(cidr string) (*SubnetInfo, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}

	info := &SubnetInfo{
		Network:   ipNet.String(),
		IPVersion: getIPVersion(ipNet.IP),
	}

	prefixLen, _ := ipNet.Mask.Size()
	info.PrefixLen = prefixLen

	if info.IPVersion == IPv4 {
		info.Mask = net.IP(ipNet.Mask).To4().String()
		total := int64(1) << uint(32-prefixLen)
		info.TotalAddresses = total
		if prefixLen <= 30 {
			info.UsableAddresses = total - 2
		} else if prefixLen == 31 {
			info.UsableAddresses = 2
		} else {
			info.UsableAddresses = 1
		}
		info.FirstIP = ipNet.IP.To4().String()
		lastIP := make(net.IP, 4)
		binary.BigEndian.PutUint32(lastIP, binary.BigEndian.Uint32(ipNet.IP.To4())+uint32(total-1))
		info.LastIP = lastIP.String()
	} else {
		total := new(big.Int).Lsh(big.NewInt(1), uint(128-prefixLen))
		info.TotalAddresses = total.Int64()
		info.UsableAddresses = info.TotalAddresses
		info.FirstIP = ipNet.IP.To16().String()
		lastIP := make(net.IP, 16)
		for i := range ipNet.IP {
			lastIP[i] = ipNet.IP[i] | ^ipNet.Mask[i]
		}
		info.LastIP = net.IP(lastIP).To16().String()
	}

	return info, nil
}

// ValidateCIDR checks if a CIDR string is valid.
func ValidateCIDR(cidr string) error {
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR notation: %w", err)
	}
	return nil
}

// IPToBytes converts an IP string to its binary representation.
func IPToBytes(ipStr string) ([]byte, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ipStr)
	}
	if v4 := ip.To4(); v4 != nil {
		return v4, nil
	}
	return ip.To16(), nil
}

// BytesToIP converts binary representation to IP string.
func BytesToIP(b []byte) string {
	return net.IP(b).String()
}

// CheckOverlap checks if two CIDR ranges overlap.
func CheckOverlap(cidr1, cidr2 string) (bool, error) {
	_, net1, err := net.ParseCIDR(cidr1)
	if err != nil {
		return false, fmt.Errorf("invalid first CIDR: %w", err)
	}
	_, net2, err := net.ParseCIDR(cidr2)
	if err != nil {
		return false, fmt.Errorf("invalid second CIDR: %w", err)
	}

	return net1.Contains(net2.IP) || net2.Contains(net1.IP), nil
}

// IsSubnetOf checks if child is a subnet of parent.
func IsSubnetOf(parent, child string) (bool, error) {
	_, parentNet, err := net.ParseCIDR(parent)
	if err != nil {
		return false, fmt.Errorf("invalid parent CIDR: %w", err)
	}
	_, childNet, err := net.ParseCIDR(child)
	if err != nil {
		return false, fmt.Errorf("invalid child CIDR: %w", err)
	}

	return parentNet.Contains(childNet.IP), nil
}

// GetFirstAvailableIP returns the first available IP in a subnet given existing IPs.
func GetFirstAvailableIP(cidr string, existingIPs []string) (string, error) {
	info, err := ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	existingSet := make(map[string]bool)
	for _, ip := range existingIPs {
		existingSet[ip] = true
	}

	ip := net.ParseIP(info.FirstIP)
	if ip == nil {
		return "", fmt.Errorf("invalid first IP")
	}

	if info.IPVersion == IPv4 && info.PrefixLen <= 30 {
		ip = ip.To4()
		incrementIP(ip)
	}

	lastIP := net.ParseIP(info.LastIP)
	for !ip.Equal(lastIP) {
		ipStr := ip.String()
		if !existingSet[ipStr] {
			return ipStr, nil
		}
		incrementIP(ip)
	}

	if !existingSet[lastIP.String()] && info.PrefixLen >= 31 {
		return lastIP.String(), nil
	}

	return "", fmt.Errorf("no available IP addresses in subnet %s", cidr)
}

// GetAvailableRanges returns available IP ranges in a subnet given existing IPs.
func GetAvailableRanges(cidr string, existingIPs []string) ([]IPRange, error) {
	info, err := ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	existingSet := make(map[string]bool)
	for _, ip := range existingIPs {
		existingSet[ip] = true
	}

	var ranges []IPRange
	ip := net.ParseIP(info.FirstIP)
	lastIP := net.ParseIP(info.LastIP)

	var rangeStart net.IP
	for !ip.Equal(lastIP) {
		ipStr := ip.String()
		if !existingSet[ipStr] {
			if rangeStart == nil {
				rangeStart = make(net.IP, len(ip))
				copy(rangeStart, ip)
			}
		} else {
			if rangeStart != nil {
				prev := make(net.IP, len(ip))
				copy(prev, ip)
				decrementIP(prev)
				ranges = append(ranges, IPRange{
					Start: rangeStart.String(),
					End:   prev.String(),
				})
				rangeStart = nil
			}
		}
		incrementIP(ip)
	}

	if !existingSet[lastIP.String()] {
		if rangeStart != nil {
			ranges = append(ranges, IPRange{
				Start: rangeStart.String(),
				End:   lastIP.String(),
			})
		} else {
			ranges = append(ranges, IPRange{
				Start: lastIP.String(),
				End:   lastIP.String(),
			})
		}
	} else if rangeStart != nil {
		prev := make(net.IP, len(lastIP))
		copy(prev, lastIP)
		decrementIP(prev)
		ranges = append(ranges, IPRange{
			Start: rangeStart.String(),
			End:   prev.String(),
		})
	}

	for i := range ranges {
		ranges[i].Count = countIPsBetween(ranges[i].Start, ranges[i].End)
	}

	return ranges, nil
}

// SplitSubnet splits a subnet into two equal halves.
func SplitSubnet(cidr string) ([]string, error) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}

	prefixLen, bits := ipNet.Mask.Size()
	if prefixLen >= bits {
		return nil, fmt.Errorf("cannot split a /%d subnet", prefixLen)
	}

	newPrefix := prefixLen + 1
	mask := net.CIDRMask(newPrefix, bits)

	// First half: same network address, longer prefix
	firstIP := ipNet.IP.Mask(mask)
	firstNet := &net.IPNet{IP: firstIP, Mask: mask}

	// Second half: flip the bit at position newPrefix-1
	secondIP := make(net.IP, len(ipNet.IP))
	copy(secondIP, ipNet.IP)
	byteIndex := (newPrefix - 1) / 8
	bitIndex := uint(7 - (newPrefix-1)%8)
	secondIP[byteIndex] |= 1 << bitIndex
	secondIPMasked := secondIP.Mask(mask)
	secondNet := &net.IPNet{IP: secondIPMasked, Mask: mask}

	return []string{firstNet.String(), secondNet.String()}, nil
}

func getIPVersion(ip net.IP) IPVersion {
	if v4 := ip.To4(); v4 != nil {
		return IPv4
	}
	return IPv6
}

func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}

func decrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]--
		if ip[i] != 0xff {
			break
		}
	}
}

func countIPsBetween(start, end string) int64 {
	s := net.ParseIP(start)
	e := net.ParseIP(end)
	if s == nil || e == nil {
		return 0
	}

	if v4s := s.To4(); v4s != nil {
		if v4e := e.To4(); v4e != nil {
			return int64(binary.BigEndian.Uint32(v4e)) - int64(binary.BigEndian.Uint32(v4s)) + 1
		}
	}

	return 1
}
