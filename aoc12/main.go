package main

import (
	"container/heap"
	"fmt"
	"strings"

	"cjting.me/aoc2022/utils"
)

type Dir int

const MAX_INT = int(^uint(0) >> 1)

const (
	Up Dir = iota
	Down
	Left
	Right
)

func (d Dir) String() string {
	return []string{"Up", "Down", "Left", "Right"}[int(d)]
}

func main() {
	stdin := utils.ReadStdin()
	m, start, end := parse(stdin)

	// 472
	// part1(m, start, end)

	// 465
	part2(m, start, end)
}

// top-left point is the origin
type Pos struct {
	x int
	y int
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Map [][]rune

func (m Map) X() int {
	return len(m[0])
}
func (m Map) Y() int {
	return len(m)
}
func (m Map) toIndex(x, y int) int {
	return y*m.X() + x
}
func (m Map) toXY(index int) (int, int) {
	y := index / m.X()
	x := index % m.X()
	return x, y
}

func (m Map) String() string {
	var lines []string
	for _, l := range m {
		lines = append(lines, string(l))
	}
	return strings.Join(lines, "\n")
}

func nextPos(dir Dir, x, y int) (int, int) {
	nextX := x
	nextY := y

	switch dir {
	case Up:
		nextY -= 1
	case Down:
		nextY += 1
	case Left:
		nextX -= 1
	case Right:
		nextX += 1
	}

	return nextX, nextY
}

// 0:up, 1:down, 2:left, 3:right
func (m Map) canGoDirection(dir Dir, x, y int) bool {
	nextX, nextY := nextPos(dir, x, y)

	if nextX < 0 || nextX >= m.X() || nextY < 0 || nextY >= m.Y() {
		return false
	}

	curElevation := m[y][x]
	nextElevation := m[nextY][nextX]

	return nextElevation <= curElevation+1
}

func parse(str string) (Map, int, int) {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	var m Map
	var startX, startY, endX, endY int

	for y, line := range lines {
		runes := []rune(line)

		sIdx := strings.Index(line, "S")
		eIdx := strings.Index(line, "E")

		if sIdx != -1 {
			startX = sIdx
			startY = y
			runes[sIdx] = 'a'
		}

		if eIdx != -1 {
			endX = eIdx
			endY = y
			runes[eIdx] = 'z'
		}

		m = append(m, runes)
	}

	return m, m.toIndex(startX, startY), m.toIndex(endX, endY)
}

type HeapItem struct {
	index int
	cost  int
}
type MinHeap []HeapItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(HeapItem))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) push(index, cost int) {
	heap.Push(h, HeapItem{index, cost})
}
func (h *MinHeap) pop() int {
	return heap.Pop(h).(HeapItem).index
}

// return -1 means no path from startIndex to endIndex
func getFewestSteps(m Map, startIndex, endIndex int) (int, map[int]int) {
	// dirjstra algorithm
	costSoFar := make(map[int]int)
	visited := make(map[int]bool)
	parent := make(map[int]int)

	frontier := &MinHeap{}
	heap.Init(frontier)

	// bootstrap
	costSoFar[startIndex] = 0
	frontier.push(startIndex, 0)

	// let's go
	for frontier.Len() > 0 {
		cur := frontier.pop()

		if cur == endIndex {
			return costSoFar[cur], parent
		}

		if visited[cur] {
			continue
		}

		visited[cur] = true

		x, y := m.toXY(cur)

		for _, dir := range []Dir{Up, Down, Left, Right} {
			if !m.canGoDirection(dir, x, y) {
				continue
			}

			nextX, nextY := nextPos(dir, x, y)
			newIndex := m.toIndex(nextX, nextY)
			newCost := costSoFar[cur] + 1

			_, ok := costSoFar[newIndex]
			if !ok || newCost < costSoFar[newIndex] {
				costSoFar[newIndex] = newCost
				frontier.push(newIndex, newCost)
				parent[newIndex] = cur
			}
		}
	}

	return -1, nil
}

func printPath(parent map[int]int, endIndex int) {
	path := []int{}

	cur := endIndex
	var ok bool

	for {
		path = append(path, cur)
		cur, ok = parent[cur]

		if !ok {
			break
		}
	}

	for i := len(path) - 1; i >= 1; i-- {
		fmt.Print(path[i], "->", path[i-1])
	}
	fmt.Println()
}

func part1(m Map, startIndex, endIndex int) {
	steps, parent := getFewestSteps(m, startIndex, endIndex)
	fmt.Println(steps)

	_ = parent
	// printPath(parent, endIndex)
}

func part2(m Map, _, endIndex int) {
	min := int(MAX_INT)

	for y := 0; y < m.Y(); y++ {
		for x := 0; x < m.X(); x++ {
			if m[y][x] == 'a' {
				result, _ := getFewestSteps(m, m.toIndex(x, y), endIndex)
				if result >= 0 && result < min {
					min = result
				}
			}
		}
	}

	fmt.Println(min)
}
