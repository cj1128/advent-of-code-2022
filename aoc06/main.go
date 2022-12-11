package main

import (
	"fmt"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := strings.TrimSpace(utils.ReadStdin())

	// 1531
	// part1(stdin)

	part2(stdin)
}

func hasDuplicate(runes []rune) bool {
	m := make(map[rune]bool)

	for _, r := range runes {
		if m[r] == true {
			return true
		}
		m[r] = true
	}

	return false
}

// return -1 if not found
func findMarkerPosition(str string) int {
	result := -1
	runes := []rune(str)

	for i := 0; i <= len(runes)-4; i++ {
		if !hasDuplicate(runes[i : i+4]) {
			return i + 4
		}
	}

	return result
}

// return -1 if not found
func findMessagePosition(str string) int {
	result := -1
	runes := []rune(str)

	for i := 0; i <= len(runes)-14; i++ {
		if !hasDuplicate(runes[i : i+14]) {
			return i + 14
		}
	}

	return result
}

func testFindMarkerPosition() {
	cases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range cases {
		result := findMarkerPosition(c.input)
		if result != c.expected {
			fmt.Printf("input '%s', got %d, expect %d \n", c.input, result, c.expected)
		}
	}
}

func testFindMessagePosition() {
	cases := []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, c := range cases {
		result := findMessagePosition(c.input)
		if result != c.expected {
			fmt.Printf("input '%s', got %d, expect %d \n", c.input, result, c.expected)
		}
	}
}

func part1(str string) {
	// testFindMarkerPosition()
	fmt.Println(findMarkerPosition(str))
}

func part2(str string) {
	// testFindMessagePosition()
	fmt.Println(findMessagePosition(str))
}
