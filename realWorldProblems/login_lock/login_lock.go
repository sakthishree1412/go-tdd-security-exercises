package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	logs := []string{
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:01Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:02:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:04:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:06:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:09:00Z",
		"user=bob ip=10.0.0.2 status=success timestamp=2025-07-10T10:01:00Z",
		"user=bob ip=10.0.0.2 status=fail timestamp=2025-07-10T10:02:00Z",
		"user=bob ip=10.0.0.2 status=fail timestamp=2025-07-10T10:04:00Z",
	}

	lockedUsers := GetLoginLock(logs)
	fmt.Println("Locked out users:", lockedUsers)
}
func GetLoginLock(logs []string) []string {
	lockeduser := make(map[string][]time.Time)
	var user string
	var status string
	var parsedTime time.Time
	for _, space := range logs {
		parts := strings.Split(space, " ")
		for _, part := range parts {
			if strings.HasPrefix(part, "user=") {
				user = strings.TrimPrefix(part, "user=")
			}
			if strings.HasPrefix(part, "status=") {
				status = strings.TrimPrefix(part, "status=")
			}
			if strings.HasPrefix(part, "timestamp=") {
				timeStr := strings.TrimPrefix(part, "timestamp=")
				parsedTime, _ = time.Parse(time.RFC3339, timeStr)
			}

		}
		if status == "fail" {
			lockeduser[user] = append(lockeduser[user], parsedTime)
		}

	}
	var lockedInUser []string
	seen := make(map[string]bool)
	for k, v := range lockeduser {

		sort.Slice(v, func(i, j int) bool {
			return v[i].Before(v[j])
		})
		i := 0
		for j := 0; j < len(v); j++ {
			if v[j].Sub(v[i]) > 10*time.Minute {
				i++
			}
			fmt.Println("i value, j value", i, j)
			if j-i+1 >= 5 && !seen[k] {
				lockedInUser = append(lockedInUser, k)
				seen[k] = true

			}
		}

	}
	return lockedInUser
}
