package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	parsed := parse(utils.ReadStdin())

	// 5393
	// part1(parsed)

	// 26712
	part2(parsed)
}

type Packet []any
type Pair [2]Packet

func parse(str string) []Pair {
	uints := strings.Split(strings.TrimSpace(str), "\n\n")

	var result []Pair

	for _, u := range uints {
		lines := strings.Split(u, "\n")
		utils.Assert(len(lines) == 2)

		pair := Pair{parsePacket(lines[0]), parsePacket(lines[1])}

		result = append(result, pair)
	}

	return result
}

func part1(pairs []Pair) {
	result := 0

	for idx, pair := range pairs {
		if cmp(pair[0], pair[1]) == 1 {
			// fmt.Println("right order", idx+1)
			result += idx + 1
		}
	}

	fmt.Println(result)
}

func part2(pairs []Pair) {
	var allPackets []Packet

	for _, p := range pairs {
		allPackets = append(allPackets, p[0])
		allPackets = append(allPackets, p[1])
	}

	d1 := parsePacket("[[2]]")
	d2 := parsePacket("[[6]]")

	allPackets = append(allPackets, d1)
	allPackets = append(allPackets, d2)

	sort.Slice(allPackets, func(i, j int) bool {
		return cmp(allPackets[i], allPackets[j]) == 1
	})

	idx1 := 0
	idx2 := 0

	for i, p := range allPackets {
		if reflect.DeepEqual(p, d1) {
			idx1 = i + 1
		}
		if reflect.DeepEqual(p, d2) {
			idx2 = i + 1
		}
	}

	fmt.Println(idx1 * idx2)
}

// 1: a < b
// 0: a == b
// -1: a > b
func cmp(a []any, b []any) int {
	minLen := utils.Min(len(a), len(b))

	for i := 0; i < minLen; i++ {
		item1 := a[i]
		item2 := b[i]

		if isNumber(item1) && isNumber(item2) {
			f1 := item1.(float64)
			f2 := item2.(float64)

			if f1 < f2 {
				return 1
			}

			if f1 > f2 {
				return -1
			}

			continue
		}

		if isNumber(item1) {
			item1 = []any{item1}
		}
		if isNumber(item2) {
			item2 = []any{item2}
		}

		r := cmp(item1.([]any), item2.([]any))

		if r == 0 {
			continue
		}

		return r
	}

	if len(a) < len(b) {
		return 1
	}

	if len(a) > len(b) {
		return -1
	}

	return 0
}

func parsePacket(str string) Packet {
	var result Packet

	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}

	return result
}

func isList(item any) bool {
	_, ok := item.([]any)
	return ok
}
func isNumber(item any) bool {
	_, ok := item.(float64)
	return ok
}
