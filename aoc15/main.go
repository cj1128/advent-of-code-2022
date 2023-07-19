package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

type Pair struct {
	sx int
	sy int
	bx int
	by int
}

func main() {
	pairs := parse(utils.ReadStdin())

	// 6078701
	// part1(pairs)

	// 12567351400528
	part2(pairs)
}

func part1(pairs []Pair) {
	targetY := 2000000
	// targetY := 10

	ranges := getRanges(pairs, targetY)

	result := 0

	seen := make(map[string]bool)
	for _, pair := range pairs {
		bx := pair.bx
		by := pair.by

		if by != targetY {
			continue
		}

		key := fmt.Sprintf("%d.%d", bx, by)
		if seen[key] {
			continue
		}

		for _, r := range ranges {
			if bx >= r.start && bx <= r.end {
				seen[key] = true
				result--
			}
		}
	}

	for _, r := range ranges {
		result += r.end - r.start + 1
	}

	fmt.Println(result)
}

func getRanges(pairs []Pair, targetY int) []Range {
	var ranges []Range

	for _, pair := range pairs {
		sx := pair.sx
		sy := pair.sy
		bx := pair.bx
		by := pair.by

		dis := abs(sx-bx) + abs(sy-by)

		minY := sy - dis
		maxY := sy + dis

		if maxY < targetY || minY > targetY {
			continue
		}

		startX := sx - dis + abs(targetY-sy)
		endX := sx + dis - abs(targetY-sy)

		ranges = append(ranges, Range{startX, endX})
	}

	return normalize((ranges))
}

type Range struct {
	start int
	end   int
}

func normalize(ranges []Range) []Range {
	if len(ranges) == 0 {
		return []Range{}
	}

	// Sort the intervals based on their start values
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	res := []Range{ranges[0]} // Initialize the result with the first interval

	for _, interval := range ranges[1:] {
		if interval.start <= res[len(res)-1].end {
			// If the current interval starts before the last interval in the result ends,
			// merge the current interval into the last interval in the result.
			if interval.end > res[len(res)-1].end {
				res[len(res)-1].end = interval.end
			}
		} else {
			// If the current interval starts after the last interval in the result ends,
			// append the current interval to the result.
			res = append(res, interval)
		}
	}

	return res
}

func part2(pairs []Pair) {
	max := 4000000

	resultX := 0
	resultY := 0

	for y := 0; y <= max; y++ {
		ranges := getRanges(pairs, y)

		if len(ranges) == 0 {
			continue
		}

		if ranges[0].start < 0 && ranges[0].end > max {
			continue
		}

		// found!
		resultX = ranges[1].start - 1
		resultY = y

		fmt.Println("found!", ranges)
		break
	}

	fmt.Println(resultX*max + resultY)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parse(str string) []Pair {
	var result []Pair

	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	regex := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		matched := regex.FindStringSubmatch(line)

		if matched == nil {
			panic(fmt.Sprintf("invalid input: %s", line))
		}

		sx := mustParseInt(matched[1])
		sy := mustParseInt(matched[2])
		bx := mustParseInt(matched[3])
		by := mustParseInt(matched[4])

		result = append(result, Pair{sx, sy, bx, by})
	}

	return result
}

func mustParseInt(str string) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		panic(fmt.Sprintf("invalid number string: %s", str))
	}
	return result
}
