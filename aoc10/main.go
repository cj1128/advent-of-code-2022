package main

import (
	"fmt"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	ins := parse(stdin)

	// 13440
	// part1(ins)

	// ###..###..####..##..###...##..####..##..
	// #..#.#..#....#.#..#.#..#.#..#....#.#..#.
	// #..#.###....#..#....#..#.#..#...#..#..#.
	// ###..#..#..#...#.##.###..####..#...####.
	// #....#..#.#....#..#.#.#..#..#.#....#..#.
	// #....###..####..###.#..#.#..#.####.#..#.
	//
	// PBZGRAZA
	part2(ins)
}

type Command int

const (
	Noop Command = iota
	AddX
)

func (c Command) String() string {
	return []string{"Noop", "AddX"}[c]
}

type Instruction struct {
	command Command
	arg     int
}

func (i Instruction) String() string {
	if i.command == Noop {
		return "[noop]"
	}

	return fmt.Sprintf("[addx %d]", i.arg)
}

func parse(str string) []Instruction {
	var result []Instruction

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		parts := strings.Split(line, " ")

		if parts[0] == "noop" {
			result = append(result, Instruction{Noop, 0})
			continue
		}

		if parts[0] == "addx" {
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				utils.Err("invalid command: %s", line)
			}

			result = append(result, Instruction{AddX, num})
			continue
		}

		utils.Err("invalid command: %s", line)
	}

	return result
}

func part1(ins []Instruction) {
	cycle := 0

	register := make(map[string]int)
	register["x"] = 1

	result := 0

	calc := func(prevX, prevCycle, cycle int) {
		for _, target := range []int{20, 60, 100, 140, 180, 220} {
			if prevCycle < target && cycle >= target {
				// fmt.Println(target, prevX)
				result += prevX * target
			}
		}
	}

	for _, in := range ins {
		// before the cycle
		prevX := register["x"]
		prevCycle := cycle

		cycle += tick(register, in)

		// after the cycle
		calc(prevX, prevCycle, cycle)
		// fmt.Println("after", cycle, register["x"])
	}

	fmt.Println(result)
}

// return cycle number that the instruction consumes
func tick(register map[string]int, in Instruction) int {
	switch in.command {
	case Noop:
		{
			return 1
		}
	case AddX:
		{
			register["x"] += in.arg
			return 2
		}
	}

	panic("unreachable!")
}

type CRT struct {
	// 40x6
	pixels [][]rune
}

func (c *CRT) draw(x, y int, char rune) {
	// fmt.Println("draw", x, y)
	if char != '#' && char != '.' {
		utils.Err("invalid char being drawn: %v", char)
	}

	c.pixels[y][x] = char
}

func (c *CRT) String() string {
	lines := make([]string, 6)

	for idx := range c.pixels {
		lines[idx] = string(c.pixels[idx])
	}

	return strings.Join(lines, "\n")
}

func newCRT() *CRT {
	pixels := make([][]rune, 6)

	for y := 0; y < 6; y++ {
		pixels[y] = make([]rune, 40)

		for idx := range pixels[y] {
			pixels[y][idx] = '_'
		}
	}

	return &CRT{pixels}
}

func part2(ins []Instruction) {
	crt := newCRT()

	// sprite middle point position
	spritePos := 1

	register := make(map[string]int)
	register["x"] = 1
	cycle := 0

	draw := func(startCycle, spritePos, delta int) {
		for d := 1; d <= delta; d++ {
			currentCycle := startCycle + d
			x := (currentCycle - 1) % 40
			y := ((currentCycle - 1) / 40) % 6

			char := '.'

			if x >= spritePos-1 && x <= spritePos+1 {
				char = '#'
			}

			crt.draw(x, y, char)
		}

	}

	for _, in := range ins {
		// before the cycle
		delta := tick(register, in)
		// after the cycle
		draw(cycle, spritePos, delta)

		spritePos = register["x"]
		cycle += delta

		// fmt.Println()
		// fmt.Println(crt)
		// fmt.Println()
	}

	fmt.Println(crt)
}
