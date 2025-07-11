package main

import (
	"reflect"
	"testing"
)

func TestDetectRateLimitViolations(t *testing.T) {
	logs := []string{
		"user=alice timestamp=2025-07-10T10:00:01Z",
		"user=alice timestamp=2025-07-10T10:00:10Z",
		"user=alice timestamp=2025-07-10T10:00:20Z",
		"user=alice timestamp=2025-07-10T10:00:30Z",
		"user=alice timestamp=2025-07-10T10:00:40Z",
		"user=alice timestamp=2025-07-10T10:00:50Z", // 6th in 60s window
		"user=bob timestamp=2025-07-10T10:05:00Z",
		"user=bob timestamp=2025-07-10T10:05:50Z",
	}

	expected := []string{"alice"}

	result := DetectRateLimitViolations(logs, 5)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
