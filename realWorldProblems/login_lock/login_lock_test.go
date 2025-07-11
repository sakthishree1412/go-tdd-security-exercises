package main

import "testing"

func TestLoginLock(t *testing.T) {
	logs := []string{
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:00:01Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:02:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:04:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:06:00Z",
		"user=alice ip=10.0.0.1 status=fail timestamp=2025-07-10T10:09:00Z",
		"user=bob ip=10.0.0.2 status=success timestamp=2025-07-10T10:01:00Z",
	}
	expected := []string{"alice"}

	result := GetLoginLock(logs)

	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
