package main

import (
	"fmt"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	cave := newCave(utils.ReadStdin())
	// pp.Println(cave)

	// 757
	// part1(cave)

	part2(cave)
}

type Status int

const (
	Air  Status = iota
	Rock        = iota
	Sand        = iota
)

// The sand is pouring into the cave from point 500,0.
type Cave struct {
	minX int
	maxX int
	// minY is always 0
	maxY int

	// 0: air
	// 1: rock
	// 2: sand
	status    map[string]Status
	sandCount int
}

type Point struct {
	x int
	y int
}

type Path []Point

func parseInput(str string) []Path {
	var paths []Path

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		var path []Point

		for _, part := range strings.Split(line, " -> ") {
			xys := strings.Split(part, ",")
			utils.Assert(len(xys) == 2)

			x, err := strconv.Atoi(xys[0])
			utils.Assert(err == nil)

			y, err := strconv.Atoi(xys[1])
			utils.Assert(err == nil)

			path = append(path, Point{x, y})
		}

		paths = append(paths, path)
	}

	return paths
}

func key(x, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

func (c *Cave) set(x, y int, status Status) {
	c.status[key(x, y)] = status
}
func (c *Cave) get(x, y int) Status {
	return c.status[key(x, y)]
}

func newCave(str string) *Cave {
	paths := parseInput(str)
	utils.Assert(len(paths) > 0 && len(paths[0]) > 0)

	result := &Cave{minX: paths[0][0].x, maxX: paths[0][0].x, maxY: paths[0][0].y, status: make(map[string]Status)}

	for _, path := range paths {
		for startIdx := 0; startIdx < len(path)-1; startIdx++ {
			start := path[startIdx]
			end := path[startIdx+1]

			// start -> end is a straight line

			// vertical line
			if start.x == end.x {
				minY, maxY := start.y, end.y
				if minY > maxY {
					minY, maxY = maxY, minY
				}

				for y := minY; y <= maxY; y++ {
					result.set(start.x, y, Rock)
				}

				if maxY > result.maxY {
					result.maxY = maxY
				}
			} else {
				minX, maxX := start.x, end.x
				if minX > maxX {
					minX, maxX = maxX, minX
				}

				for x := minX; x <= maxX; x++ {
					result.set(x, start.y, Rock)
				}

				if minX < result.minX {
					result.minX = minX
				}
				if maxX > result.maxX {
					result.maxX = maxX
				}
			}
		}
	}

	return result
}

// return false when sand can not be moved
func (c *Cave) move(x *int, y *int) bool {
	// down
	if c.isEmpty(*x, *y+1) {
		*y = *y + 1
		return true
	}

	// down left
	if c.isEmpty(*x-1, *y+1) {
		*y = *y + 1
		*x = *x - 1
		return true
	}

	// down right
	if c.isEmpty(*x+1, *y+1) {
		*y = *y + 1
		*x = *x + 1
		return true
	}

	return false
}

func (c *Cave) isEmpty(x int, y int) bool {
	return c.get(x, y) == Air
}

func part1(c *Cave) {
	// pp.Println(cave)

out:
	for {
		sx := 500
		sy := 0

		for c.move(&sx, &sy) {
			// fmt.Println("move", sx, sy)

			// falling into the endless void
			if sx < c.minX || sx > c.maxX || sy > c.maxY {
				break out
			}
		}

		c.set(sx, sy, Sand)
		c.sandCount++
	}

	fmt.Println(c.sandCount)
}

func part2(c *Cave) {
out:
	for {
		sx := 500
		sy := 0

		for c.move(&sx, &sy) {
			// fmt.Println("move", sx, sy)

			if sy == c.maxY+1 {
				break
			}
		}

		if sx == 500 && sy == 0 {
			c.sandCount++
			break out
		}

		c.set(sx, sy, Sand)
		c.sandCount++
	}

	fmt.Println(c.sandCount)
}
