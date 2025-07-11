package main

import (
	"fmt"
	"strings"
)

func DetectSuspiciousIPs(logs []string) []string {
	ips := make(map[string]int)
	var suspiousIp []string

	for _, log := range logs {
		parts := strings.Split(log, " ")
		for _, part := range parts {
			if strings.HasPrefix(part, "ip=") {
				ip := strings.TrimPrefix(part, "ip=")
				ips[ip]++
			}
		}
	}
	for ip, count := range ips {
		if count >= 4 {
			suspiousIp = append(suspiousIp, ip)
		}
	}
	fmt.Println("suspiousIp", suspiousIp)
	return suspiousIp
}

func main() {
	logs := []string{
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:01Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:15Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:30Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:55Z",
		"user=bob ip=10.0.0.2 status=success timestamp=2025-07-10T10:01:00Z",
	}

	result := DetectSuspiciousIPs(logs)
	fmt.Println("Suspicious IPs:", result)
}
