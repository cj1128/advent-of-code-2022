package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

type Elf struct {
	idx      int
	calories int
}

func parse(str string) []Elf {
	var result []Elf
	current := 0
	idx := 0

	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			result = append(result, Elf{idx: idx, calories: current})
			current = 0
			idx += 1
			continue
		}

		num, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			panic(fmt.Errorf("invalid number %s", line))
		}

		current += int(num)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].calories > result[j].calories
	})

	return result
}

func main() {
	parsed := parse(utils.ReadStdin())

	// part1
	// part1(parsed)

	// part2
	part2(parsed)
}

func part1(parsed []Elf) {
	fmt.Println(parsed[0])
}

func part2(parsed []Elf) {
	if len(parsed) < 3 {
		panic("Elves number less than 3")
	}

	sum := 0
	for _, e := range parsed[:3] {
		sum += e.calories
	}
	fmt.Println(sum)
}
