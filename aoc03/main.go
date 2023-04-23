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
	var commons [][]rune

	for _, sack := range sacks {
		common := utils.FindCommonInString([]string{sack[0], sack[1]})
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
	result := 0
	for _, g := range groups {
		common := utils.FindCommonInString(g)
		utils.Assert(len(common) == 1)
		// fmt.Println(string(common))
		result += calcScore(common)
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
