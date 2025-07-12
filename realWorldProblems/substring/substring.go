package main

import (
	"fmt"
)

func main() {
	input := "abcabcbb"
	result := LengthOfLongestSubstring(input)
	fmt.Printf("Input: %q\n", input)
	fmt.Printf("Length of longest substring without repeating characters: %d\n", result)
}

func LengthOfLongestSubstring(input string) int {
	start := 0
	seen := make(map[rune]int)
	maxLen := 0
	currentlen := 0

	for idx, i := range input {
		if lastseenidx, found := seen[i]; found && lastseenidx >= start {
			start = lastseenidx + 1
		}
		seen[i] = idx
		currentlen = idx - start + 1
		maxLen = max(maxLen, currentlen)

	}

	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
