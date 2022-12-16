package main

import (
	"fmt"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	steps := parse(stdin)

	// 6563
	// part1(steps)

	part2(steps)
}

type Step struct {
	dir Move
	num int
}

func parse(str string) []Step {
	var result []Step

	dirMapping := map[string]Move{
		"U": Up,
		"D": Down,
		"R": Right,
		"L": Left,
	}

	for _, line := range strings.Split(strings.TrimSpace(str), "\n") {
		parts := strings.Split(line, " ")

		dir, ok := dirMapping[parts[0]]
		num, err := strconv.Atoi(parts[1])

		if !ok || err != nil {
			utils.Err("invalid step string %s", line)
		}

		result = append(result, Step{
			dir,
			num,
		})
	}

	return result
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

type Move int

const (
	Left Move = iota
	Right
	Up
	Down
	TopLeft
	TopRight
	BottomLeft
	BottomRight
	None
)

func (d Move) String() string {
	return []string{"Left", "Right", "Up", "Down", "TopLeft", "TopRight", "BottomLeft", "BottomRight", "None"}[d]
}

func calcTailMove(head, tail *Pos) Move {
	xAbs := abs(head.x - tail.x)
	yAbs := abs(head.y - tail.y)

	if head.x == tail.x {
		if yAbs <= 1 {
			return None
		}

		if tail.y < head.y {
			return Up
		}

		return Down
	}

	if head.y == tail.y {
		if xAbs <= 1 {
			return None
		}

		if tail.x < head.x {
			return Right
		}

		return Left
	}

	if xAbs == 1 && yAbs == 1 {
		return None
	}

	if tail.y < head.y {
		if tail.x < head.x {
			return TopRight
		}
		return TopLeft
	}

	if tail.x < head.x {
		return BottomRight
	}
	return BottomLeft
}

func move(target *Pos, dir Move) {
	if dir == None {
		return
	}

	switch dir {
	case Up:
		{
			target.y += 1
		}
	case Down:
		{
			target.y -= 1
		}
	case Left:
		{
			target.x -= 1

		}
	case Right:
		{
			target.x += 1
		}

	case TopLeft:
		{
			target.y += 1
			target.x -= 1

		}
	case TopRight:
		{
			target.y += 1
			target.x += 1
		}

	case BottomLeft:
		{
			target.y -= 1
			target.x -= 1
		}

	case BottomRight:
		{
			target.y -= 1
			target.x += 1
		}
	}
}

// bottom-left is origin point
type Pos struct {
	x int
	y int
}

func part1(steps []Step) {
	head := &Pos{0, 0}
	tail := &Pos{0, 0}

	tailVisited := make(map[string]bool)
	tailVisited["0.0"] = true

	for _, step := range steps {
		for i := 0; i < step.num; i++ {
			move(head, step.dir)
			tm := calcTailMove(head, tail)
			move(tail, tm)
			key := fmt.Sprintf("%d.%d", tail.x, tail.y)
			tailVisited[key] = true
		}
	}

	fmt.Println(len(tailVisited))
}

func moveTails(knots []*Pos, dir Move) {
	for i := 0; i < len(knots)-1; i++ {
		tm := calcTailMove(knots[i], knots[i+1])

		if tm == None {
			return
		}

		move(knots[i+1], tm)
	}
}
func part2(steps []Step) {
	knots := make([]*Pos, 10)
	for i := 0; i < 10; i++ {
		knots[i] = &Pos{}
	}
	head := knots[0]
	tail := knots[9]

	tailVisited := make(map[string]bool)
	tailVisited["0.0"] = true

	for _, step := range steps {
		for i := 0; i < step.num; i++ {
			// fmt.Println("==== step === ")
			move(head, step.dir)
			moveTails(knots, step.dir)
			key := fmt.Sprintf("%d.%d", tail.x, tail.y)
			tailVisited[key] = true
		}
	}

	fmt.Println(len(tailVisited))
}
