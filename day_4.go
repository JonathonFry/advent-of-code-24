package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type wordSearch [][]string

func day4Part1() {
	total := 0
	search := wordSearch{}
	for _, line := range readInputAsLines(day4Input) {
		parts := strings.Split(line, "")
		search = append(search, parts)
	}

	total += search.findHorizontal()
	total += search.findVertical()

	fmt.Println(fmt.Sprintf("Day 4 Part 1: %d", total))
}

func (w wordSearch) findHorizontal() int {
	re := regexp.MustCompile(`XMAS`)

	total := 0
	for _, l := range w {
		line := strings.Join(l, "")
		matches := re.FindAllString(line, -1)
		total += len(matches)

		slices.Reverse(l)
		reversedLine := strings.Join(l, "")
		reversedMatches := re.FindAllString(reversedLine, -1)
		total += len(reversedMatches)
	}
	return total
}

func (w wordSearch) findVertical() int {
	re := regexp.MustCompile(`XMAS`)

	verticalWords := []string{}
	for i := 0; i < len(w); i++ {
		verticalWord := []string{}
		for j := 0; j < len(w); j++ {
			verticalWord = append(verticalWord, w[j][i])
		}
		verticalWords = append(verticalWords, strings.Join(verticalWord, ""))
	}

	total := 0
	for _, v := range verticalWords {
		matches := re.FindAllString(v, -1)
		total += len(matches)

		v = reverseString(v)
		reversedMatches := re.FindAllString(v, -1)
		total += len(reversedMatches)
	}

	return total
}
