package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Part1() {
	total := 0
	for _, line := range readInputAsLines(day2Input) {

		parts := strings.Split(line, " ")

		var levels []int

		for _, part := range parts {
			v, err := strconv.Atoi(part)
			if err != nil {
				panic("Error converting string to int")
			}
			levels = append(levels, v)
		}

		// check if all decreasing or all increasing
		if allAscending(levels) || allDescending(levels) {
			total += 1
		}

	}
	fmt.Println(fmt.Sprintf("Day 2 Part 1: %d", total))
}

func allAscending(input []int) bool {
	value := input[0]
	for i, v := range input {
		if i == 0 {
			continue
		}
		if v <= value {
			return false
		}
		diff := absInt(v - value)
		if diff < 1 || diff > 3 {
			return false
		}

		value = v
	}

	return true
}

func allDescending(input []int) bool {
	value := input[0]
	for i, v := range input {
		if i == 0 {
			continue
		}
		if v >= value {
			return false
		}
		diff := absInt(v - value)
		if diff < 1 || diff > 3 {
			return false
		}
		value = v
	}

	return true
}
