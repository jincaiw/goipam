package ipam

import (
	"testing"
)

func TestParseCIDR(t *testing.T) {
	tests := []struct {
		name       string
		cidr       string
		wantFirst  string
		wantLast   string
		wantTotal  int64
		wantUsable int64
		wantMask   string
		wantVer    IPVersion
	}{
		{
			name: "IPv4 /24", cidr: "192.168.1.0/24",
			wantFirst: "192.168.1.0", wantLast: "192.168.1.255",
			wantTotal: 256, wantUsable: 254, wantMask: "255.255.255.0", wantVer: IPv4,
		},
		{
			name: "IPv4 /30", cidr: "10.0.0.0/30",
			wantFirst: "10.0.0.0", wantLast: "10.0.0.3",
			wantTotal: 4, wantUsable: 2, wantMask: "255.255.255.252", wantVer: IPv4,
		},
		{
			name: "IPv4 /31", cidr: "10.0.0.0/31",
			wantFirst: "10.0.0.0", wantLast: "10.0.0.1",
			wantTotal: 2, wantUsable: 2, wantVer: IPv4,
		},
		{
			name: "IPv4 /32", cidr: "10.0.0.1/32",
			wantFirst: "10.0.0.1", wantLast: "10.0.0.1",
			wantTotal: 1, wantUsable: 1, wantVer: IPv4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := ParseCIDR(tt.cidr)
			if err != nil {
				t.Fatalf("ParseCIDR() error = %v", err)
			}
			if info.FirstIP != tt.wantFirst {
				t.Errorf("FirstIP = %v, want %v", info.FirstIP, tt.wantFirst)
			}
			if info.LastIP != tt.wantLast {
				t.Errorf("LastIP = %v, want %v", info.LastIP, tt.wantLast)
			}
			if info.TotalAddresses != tt.wantTotal {
				t.Errorf("TotalAddresses = %v, want %v", info.TotalAddresses, tt.wantTotal)
			}
			if info.UsableAddresses != tt.wantUsable {
				t.Errorf("UsableAddresses = %v, want %v", info.UsableAddresses, tt.wantUsable)
			}
			if info.IPVersion != tt.wantVer {
				t.Errorf("IPVersion = %v, want %v", info.IPVersion, tt.wantVer)
			}
			if tt.wantMask != "" && info.Mask != tt.wantMask {
				t.Errorf("Mask = %v, want %v", info.Mask, tt.wantMask)
			}
		})
	}
}

func TestValidateCIDR(t *testing.T) {
	tests := []struct {
		cidr string
		want bool
	}{
		{"192.168.1.0/24", true},
		{"10.0.0.0/8", true},
		{"invalid", false},
		{"192.168.1.0/33", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			err := ValidateCIDR(tt.cidr)
			if (err == nil) != tt.want {
				t.Errorf("ValidateCIDR(%q) = %v, want valid=%v", tt.cidr, err, tt.want)
			}
		})
	}
}

func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		name  string
		cidr1 string
		cidr2 string
		want  bool
	}{
		{"overlapping", "192.168.1.0/24", "192.168.1.128/25", true},
		{"non-overlapping", "192.168.1.0/24", "192.168.2.0/24", false},
		{"subnet contains", "10.0.0.0/8", "10.1.0.0/16", true},
		{"adjacent", "192.168.0.0/24", "192.168.1.0/24", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckOverlap(tt.cidr1, tt.cidr2)
			if err != nil {
				t.Fatalf("CheckOverlap() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("CheckOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubnetOf(t *testing.T) {
	tests := []struct {
		name   string
		parent string
		child  string
		want   bool
	}{
		{"child of", "10.0.0.0/8", "10.1.0.0/16", true},
		{"not child of", "10.0.0.0/8", "192.168.0.0/16", false},
		{"same network", "10.0.0.0/8", "10.0.0.0/8", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsSubnetOf(tt.parent, tt.child)
			if err != nil {
				t.Fatalf("IsSubnetOf() error = %v", err)
			}
			if got != tt.want {
				t.Errorf("IsSubnetOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitSubnet(t *testing.T) {
	tests := []struct {
		cidr  string
		want1 string
		want2 string
	}{
		{"192.168.1.0/24", "192.168.1.0/25", "192.168.1.128/25"},
		{"10.0.0.0/30", "10.0.0.0/31", "10.0.0.2/31"},
	}

	for _, tt := range tests {
		t.Run(tt.cidr, func(t *testing.T) {
			got, err := SplitSubnet(tt.cidr)
			if err != nil {
				t.Fatalf("SplitSubnet() error = %v", err)
			}
			if len(got) != 2 {
				t.Fatalf("SplitSubnet() returned %d parts, want 2", len(got))
			}
			if got[0] != tt.want1 || got[1] != tt.want2 {
				t.Errorf("SplitSubnet() = %v, want [%v, %v]", got, tt.want1, tt.want2)
			}
		})
	}
}

func TestGetFirstAvailableIP(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		existing []string
		want     string
		wantErr  bool
	}{
		{
			name:     "empty subnet /24",
			cidr:     "192.168.1.0/24",
			existing: []string{},
			want:     "192.168.1.1",
		},
		{
			name:     "some occupied /24",
			cidr:     "192.168.1.0/24",
			existing: []string{"192.168.1.1", "192.168.1.2"},
			want:     "192.168.1.3",
		},
		{
			name:     "full /30",
			cidr:     "10.0.0.0/30",
			existing: []string{"10.0.0.1", "10.0.0.2"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFirstAvailableIP(tt.cidr, tt.existing)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetFirstAvailableIP() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("GetFirstAvailableIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPToBytes(t *testing.T) {
	tests := []struct {
		ip   string
		want int
	}{
		{"192.168.1.1", 4},
		{"::1", 16},
	}

	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			got, err := IPToBytes(tt.ip)
			if err != nil {
				t.Fatalf("IPToBytes() error = %v", err)
			}
			if len(got) != tt.want {
				t.Errorf("IPToBytes() length = %v, want %v", len(got), tt.want)
			}
		})
	}
}
