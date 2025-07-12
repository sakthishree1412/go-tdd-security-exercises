package main

import (
	"reflect"
	"testing"
)

func TestAbuseDetection(t *testing.T) {
	tests := []struct {
		ips      []string
		k        int
		expected []string
	}{
		{
			ips: []string{
				"192.168.1.1", "10.0.0.5", "192.168.1.1",
				"10.0.0.5", "10.0.0.5", "172.16.0.2",
			},
			k:        2,
			expected: []string{"10.0.0.5", "192.168.1.1"},
		},
		{
			ips:      []string{"1.1.1.1", "2.2.2.2", "1.1.1.1"},
			k:        1,
			expected: []string{"1.1.1.1"},
		},
		{
			ips:      []string{"3.3.3.3"},
			k:        1,
			expected: []string{"3.3.3.3"},
		},
	}
	for _, tt := range tests {
		result := AbuseDetection(tt.ips, tt.k)

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Test failed")
		}
	}
}
