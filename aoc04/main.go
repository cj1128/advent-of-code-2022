package main

import (
	"fmt"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	parsed := parse(stdin)

	// 433
	// part1(parsed)

	part2(parsed)
}

// inclusive
type Range struct {
	start int
	end   int
}

type Pair [2]Range

func parse(str string) []Pair {
	var result []Pair

	genRange := func(str string) Range {
		parts := strings.Split(str, "-")
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid range str %s", str))
		}

		start, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			panic(err)
		}

		end, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			panic(err)
		}

		return Range{start: int(start), end: int(end)}
	}

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		parts := strings.Split(line, ",")

		result = append(result, Pair{genRange(parts[0]), genRange(parts[1])})
	}

	return result
}

func part1(pairs []Pair) {
	result := 0

	for _, pair := range pairs {
		r1 := pair[0]
		r2 := pair[1]

		// one fully contains the other
		if (r1.start <= r2.start && r1.end >= r2.end) || (r1.start >= r2.start && r1.end <= r2.end) {
			result += 1
		}
	}

	fmt.Println(result)
}

func part2(pairs []Pair) {
	result := 0

	for _, pair := range pairs {
		r1 := pair[0]
		r2 := pair[1]

		// overlap
		if !((r1.start > r2.end) || (r1.end < r2.start)) {
			result += 1
		}
	}

	fmt.Println(result)
}
