package main

import (
	"fmt"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	trees := parse(stdin)

	Y := len(trees)

	if Y == 0 {
		utils.Err("invalid input")
	}

	X := len(trees[0])

	// 1851
	// part1(trees, X, Y)

	part2(trees, X, Y)
}

func parse(str string) [][]int {
	var result [][]int

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		var row []int

		for _, numStr := range strings.Split(line, "") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				utils.Err("invalid number string %s:%v", numStr, err)
			}

			row = append(row, num)
		}

		result = append(result, row)
	}

	return result
}

type Dir int

const (
	Up Dir = iota
	Down
	Left
	Right
)

func checkVisible(trees [][]int, X, Y, x, y int) bool {
	// in the edge
	if x == 0 || x == X-1 || y == 0 || y == Y-1 {
		return true
	}

	height := trees[y][x]

	checkDirection := func(dir Dir) bool {
		switch dir {
		case Up:
			{
				for ty := y - 1; ty >= 0; ty-- {
					if trees[ty][x] >= height {
						return false
					}
				}

				return true
			}
		case Down:
			{
				for ty := y + 1; ty < Y; ty++ {
					if trees[ty][x] >= height {
						return false
					}
				}

				return true
			}
		case Left:
			{
				for tx := x - 1; tx >= 0; tx-- {
					if trees[y][tx] >= height {
						return false
					}
				}

				return true
			}
		case Right:
			{
				for tx := x + 1; tx < X; tx++ {
					if trees[y][tx] >= height {
						return false
					}
				}

				return true
			}
		}

		panic("unreachable!")
	}

	if checkDirection(Up) {
		return true
	}

	if checkDirection(Down) {
		return true
	}

	if checkDirection(Left) {
		return true
	}

	if checkDirection(Right) {
		return true
	}

	return false
}

func part1(trees [][]int, X, Y int) {
	result := 0

	for y := 0; y < Y; y += 1 {
		for x := 0; x < X; x += 1 {
			if checkVisible(trees, X, Y, x, y) {
				// fmt.Println("visible", x, y)
				result += 1
			}
		}
	}

	fmt.Println(result)
}

func part2(trees [][]int, X, Y int) {
	maxScore := 0

	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			score := calcScenicScore(trees, X, Y, x, y)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}

func calcScenicScore(trees [][]int, X, Y, x, y int) int {
	if x == 0 || x == X-1 || y == 0 || y == Y-1 {
		return 0
	}

	up := calcCanSee(trees, X, Y, x, y, Up)
	down := calcCanSee(trees, X, Y, x, y, Down)
	left := calcCanSee(trees, X, Y, x, y, Left)
	right := calcCanSee(trees, X, Y, x, y, Right)

	return up * down * left * right
}

// return how many trees target poition can see
// at least one
func calcCanSee(trees [][]int, X, Y, x, y int, dir Dir) int {
	height := trees[y][x]

	result := 0

	switch dir {
	case Up:
		{
			for ty := y - 1; ty >= 0; ty-- {
				result++
				if trees[ty][x] >= height {
					break
				}
			}
		}
	case Down:
		{
			for ty := y + 1; ty < Y; ty++ {
				result++
				if trees[ty][x] >= height {
					break
				}
			}
		}
	case Left:
		{
			for tx := x - 1; tx >= 0; tx-- {
				result++
				if trees[y][tx] >= height {
					break
				}
			}
		}
	case Right:
		{
			for tx := x + 1; tx < X; tx++ {
				result++
				if trees[y][tx] >= height {
					break
				}
			}
		}
	}

	return result
}
