package scan

import (
	"context"
	"fmt"
	"net"
	"time"
)

// ProbeResult represents the result of probing a single IP.
type ProbeResult struct {
	IP      string
	Alive   bool
	Message string
}

// PingProbe does a TCP connect probe on common ports.
func PingProbe(ctx context.Context, ip string, timeout time.Duration) ProbeResult {
	ports := []int{22, 80, 443, 3389}
	for _, port := range ports {
		addr := fmt.Sprintf("%s:%d", ip, port)
		dialer := net.Dialer{Timeout: timeout}
		conn, err := dialer.DialContext(ctx, "tcp", addr)
		if err == nil {
			conn.Close()
			return ProbeResult{IP: ip, Alive: true, Message: fmt.Sprintf("tcp:%d", port)}
		}
	}
	return ProbeResult{IP: ip, Alive: false, Message: "no response"}
}
