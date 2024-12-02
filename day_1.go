package main

import (
    "fmt"
    "regexp"
    "sort"
    "strconv"
)

func day1Part1() {
    // group into two lists and parse as numbers
    a, b := []int{}, []int{}
    for _, line := range readInputAsLines(day1Input) {
        if line == "" || line == "\n" {
            continue
        }
        re := regexp.MustCompile(`\d+`)

        // Find all matches in the string
        numbers := re.FindAllString(line, -1)

        first, err := strconv.Atoi(numbers[0])
        if err != nil {
            println(fmt.Sprintf("error: %v", err))
        }
        second, err := strconv.Atoi(numbers[1])
        if err != nil {
            println(fmt.Sprintf("error: %v", err))
        }
        a = append(a, first)
        b = append(b, second)
    }

    // sort the lists from low to high
    sort.Ints(a)
    sort.Ints(b)

    // iterate through the list
    total := 0
    for i := range a {
        first, second := a[i], b[i]
        // find the difference between the two numbers (abs)
        value := absInt(first - second)
        // add up all the differences
        total += value
    }

    fmt.Println(fmt.Sprintf("Day 1 Part 1: %d", total))
}

func day1Part2() {
    // group into two lists and parse as numbers
    a, b := []int{}, []int{}
    for _, line := range readInputAsLines(day1Input) {
        if line == "" || line == "\n" {
            continue
        }
        re := regexp.MustCompile(`\d+`)

        // Find all matches in the string
        numbers := re.FindAllString(line, -1)

        first, err := strconv.Atoi(numbers[0])
        if err != nil {
            println(fmt.Sprintf("error: %v", err))
        }
        second, err := strconv.Atoi(numbers[1])
        if err != nil {
            println(fmt.Sprintf("error: %v", err))
        }
        a = append(a, first)
        b = append(b, second)
    }

    // sort the lists from low to high
    sort.Ints(a)
    sort.Ints(b)

    // Similarity score

    // iterate through second slice to transform into map and iterate int
    bMap := map[int]int{}
    for _, v := range b {
        value, ok := bMap[v]
        // First time seeing this number lets initialise to 1
        if !ok {
            bMap[v] = 1
            continue
        }
        // count number of times number in first slice appears in second slice
        bMap[v] = value + 1
    }

    total := 0
    for _, v := range a {
        value, ok := bMap[v]
        if !ok {
            continue
        }
        total += value * v
    }

    fmt.Println(fmt.Sprintf("Day 1 Part 2: %d", total))
}
