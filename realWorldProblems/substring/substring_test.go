package main

import "testing"

func TestSubString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"au", 2},
		{"dvdf", 3},
	}
	for _, test := range tests {
		result := LengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %d but got %d", test.input, test.expected, result)
		}
	}
}
