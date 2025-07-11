package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func DetectRateLimitViolations(logs []string, a int) []string {
	userWithtimestamp := make(map[string][]time.Time)
	var user string
	var parsedTime time.Time
	for _, parts := range logs {
		part := strings.Split(parts, " ")
		for _, par := range part {
			if strings.HasPrefix(par, "user=") {
				user = strings.TrimPrefix(par, "user=")
			}
			if strings.HasPrefix(par, "timestamp=") {
				TimeStr := strings.TrimPrefix(par, "timestamp=")
				Time, _ := time.Parse(time.RFC3339, TimeStr)
				parsedTime = Time

			}
			userWithtimestamp[user] = append(userWithtimestamp[user], parsedTime)
			fmt.Println("userWithtimestamp", userWithtimestamp)
		}
	}
	var suspicious []string
	for k, v := range userWithtimestamp {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Before(v[j])
		})
		i := 0
		for j := 0; j < len(v); j++ {
			if v[j].Sub(v[i]) > 60*time.Second {
				i++
			}
			if j-i+1 > a {
				fmt.Println("inside ")
				suspicious = append(suspicious, k)
				break
			}

		}
	}
	return suspicious
}
func main() {
	logs := []string{
		"user=alice timestamp=2025-07-10T10:00:01Z",
		"user=alice timestamp=2025-07-10T10:00:10Z",
		"user=alice timestamp=2025-07-10T10:00:20Z",
		"user=alice timestamp=2025-07-10T10:00:30Z",
		"user=alice timestamp=2025-07-10T10:00:40Z",
		"user=alice timestamp=2025-07-10T10:00:50Z",
		"user=bob timestamp=2025-07-10T10:05:00Z",
		"user=bob timestamp=2025-07-10T10:05:50Z",
	}

	result := DetectRateLimitViolations(logs, 5)
	fmt.Println("Rate limit violators:", result)
}
