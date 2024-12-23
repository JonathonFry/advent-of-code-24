package main

import "strings"

func readInputAsLines(input string) []string {
	trimmed := strings.TrimSuffix(input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	return strings.Split(trimmed, "\n")
}

func readInputAsLine(input string) string {
	trimmed := strings.TrimSuffix(input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")
	return trimmed
}

func reverseString(s string) string {
	// Convert the string to a byte slice
	bytes := []byte(s)

	// Reverse the byte slice
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// Convert the byte slice back to a string
	return string(bytes)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
