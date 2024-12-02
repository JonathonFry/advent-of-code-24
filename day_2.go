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

func day2Part2() {
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

		// Check if would be safe if removing the level that makes it unsafe
		if variationsMatch(levels) {
			total += 1
		}

	}
	fmt.Println(fmt.Sprintf("Day 2 Part 2: %d", total))
}

func variationsMatch(levels []int) bool {
	// generate the variations of each of the slice (removing 1 of each)
	// check if any of these are valid
	// Exit early if the original input matches
	if allAscending(levels) || allDescending(levels) {
		return true
	}

	variations := [][]int{}
	for i := range levels {
		if i == len(levels) {
			continue
		}
		removed := make([]int, len(levels))
		copy(removed, levels)

		removed = append(removed[:i], removed[i+1:]...)
		variations = append(variations, removed)
	}

	// check if it matches for a sub variation
	for _, v := range variations {
		if allAscending(v) || allDescending(v) {
			return true
		}
	}

	return false
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
