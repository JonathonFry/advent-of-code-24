package main

import "strings"

func readInputAsLines(input string) []string {
	trimmed := strings.TrimSuffix(input, "\n")
	trimmed = strings.TrimPrefix(trimmed, "\n")

	return strings.Split(trimmed, "\n")
}


func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}