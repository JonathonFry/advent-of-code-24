package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day3Part1() {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	for _, line := range readInputAsLines(day3Input) {

		total := 0
		matches := re.FindAllStringSubmatch(line, -1)
		if matches != nil {
			for _, match := range matches {
				a, err := strconv.Atoi(match[1])
				if err != nil {
					panic("failed to parse number")
				}
				b, err := strconv.Atoi(match[2])
				if err != nil {
					panic("failed to parse number")
				}

				total += a * b
			}
		} else {
			fmt.Println("No match found.")
		}
		fmt.Println(fmt.Sprintf("Day 3 Part 1: %d", total))
	}
}

func day3Part2() {
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)

	for _, line := range readInputAsLines(day3Input) {
		// delete everything between the don't() and a do()
		sanitzer := regexp.MustCompile(`don't\(\).*?do\(\)`)
		input := sanitzer.ReplaceAllString(line, "")

		total := 0
		matches := re.FindAllStringSubmatch(input, -1)
		if matches != nil {
			for _, match := range matches {
				a, err := strconv.Atoi(match[1])
				if err != nil {
					panic("failed to parse number")
				}
				b, err := strconv.Atoi(match[2])
				if err != nil {
					panic("failed to parse number")
				}

				total += a * b
			}
		} else {
			fmt.Println("No match found.")
		}
		fmt.Println(fmt.Sprintf("Day 3 Part 2: %d", total))
	}
}
