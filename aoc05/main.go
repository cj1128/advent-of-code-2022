package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	initStack, steps := parse(stdin)

	// VCTFTJQCG
	// part1(initStack, steps)

	part2(initStack, steps)
}

type Stack []byte // from bottom to top

type Step struct {
	start int
	end   int
	count int
}

func parseInitStacks(str string) []Stack {
	lines := strings.Split(str, "\n")

	lastLine := lines[len(lines)-1]
	lastLineParts := strings.Split(strings.TrimSpace(lastLine), " ")

	count, err := strconv.Atoi(lastLineParts[len(lastLineParts)-1])
	if err != nil {
		panic(fmt.Errorf("invalid last line of init stack %s: %v", lastLineParts, err))
	}

	var result []Stack

	for stackNo := 0; stackNo < count; stackNo++ {
		var stack []byte

		for j := len(lines) - 2; j >= 0; j-- {
			line := lines[j]
			// i=0, idx=1
			// i=1, idx=5
			// i=2, idx=9
			idx := stackNo*4 + 1

			if len(line) > idx && line[idx] != ' ' {
				stack = append(stack, line[idx])
			}
		}

		result = append(result, stack)
	}

	return result
}

func printStacks(msg string, stacks []Stack) {
	fmt.Print(msg)
	for idx, stack := range stacks {
		fmt.Println(idx, string(stack))
	}
}

func parse(str string) ([]Stack, []Step) {
	str = strings.Trim(str, "\n")
	parts := strings.Split(str, "\n\n")

	if len(parts) != 2 {
		panic("invalid input")
	}

	initStacks := parseInitStacks(parts[0])

	var steps []Step

	reg := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for _, line := range strings.Split(parts[1], "\n") {
		matches := reg.FindStringSubmatch(line)

		if len(matches) != 4 {
			panic(fmt.Errorf("invalid step: %v", line))
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		start, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		// NOTE: change index to be zero based
		steps = append(steps, Step{start: start - 1, end: end - 1, count: count})
	}

	return initStacks, steps
}

func part1(stacks []Stack, steps []Step) {
	doStep := func(stacks []Stack, step Step) {
		start := stacks[step.start]
		moved := start[len(start)-step.count:]

		stacks[step.start] = start[0 : len(start)-step.count]

		for i := len(moved) - 1; i >= 0; i-- {
			stacks[step.end] = append(stacks[step.end], moved[i])
		}
	}

	// printStacks("init", stacks)

	for _, step := range steps {
		doStep(stacks, step)

		// fmt.Println("step", step)
		// printStacks("after step", stacks)
		// fmt.Println()
	}

	var top []byte

	for _, stack := range stacks {
		top = append(top, stack[len(stack)-1])
	}

	fmt.Println(string(top))
}

func part2(stacks []Stack, steps []Step) {
	// printStacks("init", stacks)

	doStep := func(stacks []Stack, step Step) {
		start := stacks[step.start]
		moved := start[len(start)-step.count:]

		stacks[step.start] = start[0 : len(start)-step.count]
		stacks[step.end] = append(stacks[step.end], moved...)
	}

	for _, step := range steps {
		doStep(stacks, step)

		// fmt.Println("step", step)
		// printStacks("after step", stacks)
		// fmt.Println()
	}

	var top []byte

	for _, stack := range stacks {
		top = append(top, stack[len(stack)-1])
	}

	fmt.Println(string(top))
}
