package main

import (
	"reflect"
	"testing"
)

func TestDetectSuspiciousIPs(t *testing.T) {
	logs := []string{
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:01Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:15Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:30Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:55Z",
		"user=bob ip=10.0.0.2 status=success timestamp=2025-07-10T10:01:00Z",
	}
	expected := []string{"10.0.0.1"}

	result := DetectSuspiciousIPs(logs)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
