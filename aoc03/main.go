package main

import (
	"fmt"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()

	// 8105
	// part1(parse(stdin))

	// 2363
	part2(parse2(stdin))
}

type Rucksack [2]string

func parse(str string) []Rucksack {
	var result []Rucksack

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		haflLength := len(line) / 2
		result = append(result, Rucksack{line[0:haflLength], line[haflLength:]})
	}

	return result
}

func part1(sacks []Rucksack) {

	findCommon := func(str1, str2 string) []rune {
		result := make(map[rune]bool)

		m := make(map[rune]bool)

		for _, r := range str1 {
			m[r] = true
		}

		for _, r := range str2 {
			if m[r] == true {
				result[r] = true
			}
		}

		var slice []rune

		for k := range result {
			slice = append(slice, k)
		}

		return slice
	}

	var commons [][]rune

	for _, sack := range sacks {
		common := findCommon(sack[0], sack[1])
		commons = append(commons, common)
	}

	// fmt.Println("commons")
	// for idx, c := range commons {
	// 	fmt.Println(idx, string(c))
	// }

	result := 0
	for _, common := range commons {
		result += calcScore(common)
	}

	fmt.Println(result)
}

type Group []string

func parse2(str string) []Group {
	var result []Group
	var group []string

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		group = append(group, line)
		if len(group) == 3 {
			result = append(result, group)
			group = make([]string, 0, len(group))
		}
	}

	// at the end, group must be length 0
	if len(group) != 0 {
		panic("invalid input")
	}

	return result
}

func part2(groups []Group) {
	findCommon := func(group Group) rune {
		m := make(map[rune][]int)

		appendIfNotExists := func(s []int, num int) []int {
			found := false
			result := s

			for _, v := range s {
				if v == num {
					found = true
				}
			}

			if !found {
				result = append(result, num)
			}

			return result
		}

		for lineNo, line := range group {
			for _, r := range line {
				m[r] = appendIfNotExists(m[r], lineNo)
			}
		}

		target := '-'

		for k, v := range m {
			if len(v) == 3 {
				target = k
			}
		}
		// fmt.Println(m)

		if target == '-' {
			panic(fmt.Errorf("could not find common in group"))
		}

		return target
	}

	result := 0
	for _, g := range groups {
		common := findCommon(g)
		// fmt.Println(string(common))
		result += calcScore([]rune{common})
	}

	fmt.Println(result)
}

func calcScore(common []rune) int {
	result := 0

	for _, rune := range common {
		switch {
		case rune >= 'a' && rune <= 'z':
			{
				result += int(rune - 'a' + 1)
			}

		case rune >= 'A' && rune <= 'Z':
			{
				result += int(rune - 'A' + 27)
			}
		default:
			{
				panic(fmt.Errorf("invalid common rune %v", rune))
			}
		}
	}

	return result
}
