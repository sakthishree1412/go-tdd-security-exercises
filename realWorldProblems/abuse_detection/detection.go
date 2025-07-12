package main

import (
	"fmt"
	"sort"
)

func main() {
	ips := []string{
		"1.1.1.1", "2.2.2.2", "1.1.1.1",
	}
	s := 1

	result := AbuseDetection(ips, s)
	fmt.Println("Top", s, "frequent IPs:", result)
}
func AbuseDetection(ips []string, s int) []string {
	seen := make(map[string]int)
	abuseIps := []string{}

	for _, i := range ips {
		seen[i]++
	}

	// converting map to slice as map is unordered we cant sort the map
	mapToSlice := []struct {
		ip    string
		count int
	}{}
	for ip, count := range seen {
		mapToSlice = append(mapToSlice, struct {
			ip    string
			count int
		}{ip, count})
	}
	sort.Slice(mapToSlice, func(i, j int) bool {
		return mapToSlice[i].count > mapToSlice[j].count
	})
	fmt.Println("map to slice", mapToSlice)
	for i := 0; i < s; i++ {
		abuseIps = append(abuseIps, mapToSlice[i].ip)
	}

	return abuseIps
}
