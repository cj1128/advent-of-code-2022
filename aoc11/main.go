package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	monkeys := parse(stdin)

	// 118674
	// part1(monkeys)

	// 32333418600
	part2(monkeys)
}

type OpType int

func (ot OpType) String() string {
	return []string{"Add", "Mul", "Square"}[ot]
}

const (
	Add OpType = iota
	Mul
	Square
)

type Monkey struct {
	no       int
	items    []int   // used for part1
	newItems []*Item // used for part2

	op     func(int) int
	opType OpType
	opArg  int
	opRaw  string

	test    int // always a prime number
	testRaw string

	ifTrue  int
	ifFalse int

	inspectedTimes int
}

type Item struct {
	mods map[int]int
}

func newItem(divisibles []int, value int) *Item {
	result := &Item{}
	result.mods = make(map[int]int)

	for _, d := range divisibles {
		result.mods[d] = value % d
	}

	return result
}
func (i *Item) update(m *Monkey) {
	switch m.opType {
	case Add:
		for k, v := range i.mods {
			i.mods[k] = (v + m.opArg) % k
		}
	case Mul:
		for k, v := range i.mods {
			i.mods[k] = (v * m.opArg) % k
		}
	case Square:
		for k, v := range i.mods {
			i.mods[k] = (v * v) % k
		}
	}
}
func (i *Item) divide(val int) bool {
	return i.mods[val] == 0
}

func (m *Monkey) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("Moneky %d:", m.no))
	lines = append(lines, fmt.Sprintf("  Items: %v", m.items))
	lines = append(lines, fmt.Sprintf("  Operation: %s", m.opRaw))
	lines = append(lines, fmt.Sprintf("  OpType: %s", m.opType))
	lines = append(lines, fmt.Sprintf("  OpArg: %d", m.opArg))
	lines = append(lines, fmt.Sprintf("  Test: %d", m.test))
	lines = append(lines, fmt.Sprintf("    If true: %d", m.ifTrue))
	lines = append(lines, fmt.Sprintf("    If fasle: %d", m.ifFalse))

	return strings.Join(lines, "\n")
}

func parseMonkey(idx int, _lines []string) *Monkey {
	result := &Monkey{}

	lines := make([]string, len(_lines))
	for idx := range lines {
		lines[idx] = strings.TrimSpace(_lines[idx])
	}

	if len(lines) < 6 {
		utils.Err("invalid monkey input: %s", lines)
	}

	// no
	{
		reg := regexp.MustCompile(`Monkey (\d+):`)
		matches := reg.FindStringSubmatch(lines[0])

		result.no = utils.MustAtoi(matches[1])
		if result.no != idx {
			utils.Err("monkey index and no not matched, idx: %d, no: %d", idx, result.no)
		}
	}

	// starting items
	{
		prefix := "Starting items: "
		itemStrs := strings.Split(lines[1][len(prefix):], ", ")

		for _, str := range itemStrs {
			result.items = append(result.items, utils.MustAtoi(str))
		}
	}

	// operation
	// "Operation: new = old * 19"
	{
		reg := regexp.MustCompile(`Operation: new = old (.) (.+)`)
		matches := reg.FindStringSubmatch(lines[2])

		op := matches[1]
		argStr := matches[2]
		argNum := 0

		if argStr != "old" {
			argNum = utils.MustAtoi(argStr)
		}
		result.opArg = argNum
		result.opRaw = lines[2]
		if op == "+" {
			result.opType = Add
		} else {
			if argStr == "old" {
				result.opType = Square
			} else {
				result.opType = Mul
			}
		}

		result.op = func(p int) int {
			switch result.opType {
			case Add:
				return p + argNum
			case Mul:
				return p * argNum
			case Square:
				return p * p
			}
			panic("invalid operation found: " + lines[2])
		}
	}

	// test
	// "Test: divisible by 19"
	{
		reg := regexp.MustCompile(`Test: divisible by (\d+)`)
		matches := reg.FindStringSubmatch(lines[3])

		arg := utils.MustAtoi(matches[1])

		result.testRaw = lines[3]
		result.test = arg
	}

	// if true
	// "If true: throw to monkey 2"
	{
		line := strings.TrimSpace(lines[4])
		reg := regexp.MustCompile(`If true: throw to monkey (\d+)`)
		matches := reg.FindStringSubmatch(line)
		result.ifTrue = utils.MustAtoi(matches[1])
	}

	// if false
	// "If false: throw to monkey 0"
	{
		line := strings.TrimSpace(lines[5])
		reg := regexp.MustCompile(`If false: throw to monkey (\d+)`)
		matches := reg.FindStringSubmatch(line)
		result.ifFalse = utils.MustAtoi(matches[1])
	}

	return result
}

func parse(str string) []*Monkey {
	parts := strings.Split(strings.TrimSpace(str), "\n\n")
	var monkeys []*Monkey

	for idx, part := range parts {
		lines := strings.Split(part, "\n")
		monkeys = append(monkeys, parseMonkey(idx, lines))
	}

	var divisibles []int
	for _, m := range monkeys {
		divisibles = append(divisibles, m.test)
	}

	for _, m := range monkeys {
		for _, i := range m.items {
			m.newItems = append(m.newItems, newItem(divisibles, i))
		}
	}

	return monkeys
}

func printMonkeys(monkeys []*Monkey) {
	for _, m := range monkeys {
		fmt.Println(m)
	}
}
func printMonkeyItems(monkeys []*Monkey) {
	for _, monkey := range monkeys {
		fmt.Println(monkey.no, monkey.items)
	}
}

// we can use index to find monkeys, it's guaranteed
// to be matched with its no
func part1(monkeys []*Monkey) {
	// printMoneksy(monkeys)

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			turn(monkey, monkeys)
		}

		// fmt.Println("round", i+1)
		// printMonkeyItems(monkeys)
	}

	fmt.Println(calcMonkeyBusiness(monkeys))
}

func turn(monkey *Monkey, monkeys []*Monkey) {
	for _, level := range monkey.items {
		monkey.inspectedTimes++
		level = monkey.op(level)
		level = level / 3

		targetMonkey := monkeys[monkey.ifFalse]
		if level%monkey.test == 0 {
			targetMonkey = monkeys[monkey.ifTrue]
		}
		// fmt.Println("throw", monkey.no, targetMonkey.no, level)

		targetMonkey.items = append(targetMonkey.items, level)
	}

	monkey.items = monkey.items[:0]
}

func turn2(monkey *Monkey, monkeys []*Monkey) {
	for _, item := range monkey.newItems {
		monkey.inspectedTimes++
		item.update(monkey)

		targetMonkey := monkeys[monkey.ifFalse]
		if item.divide(monkey.test) {
			targetMonkey = monkeys[monkey.ifTrue]
		}
		// fmt.Println("throw", monkey.no, targetMonkey.no, level)

		targetMonkey.newItems = append(targetMonkey.newItems, item)
	}
	monkey.newItems = monkey.newItems[:0]
}

func calcMonkeyBusiness(monkeys []*Monkey) int {
	times := make([]int, len(monkeys))

	for idx, mon := range monkeys {
		times[idx] = mon.inspectedTimes
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i] > times[j]
	})

	return times[0] * times[1]
}

func part2(monkeys []*Monkey) {
	// printMonkeys(monkeys)

	for i := 1; i <= 10_000; i++ {
		for _, monkey := range monkeys {
			turn2(monkey, monkeys)
		}

		// if i == 1 || i == 20 || i == 1000 || i == 2000 {
		// 	fmt.Println("round", i)
		// 	for _, mon := range monkeys {
		// 		fmt.Print(mon.inspectedTimes)
		// 		fmt.Print(" ")
		// 		// fmt.Println(mon.items)
		// 	}
		// 	fmt.Println()
		// }
	}

	fmt.Println(calcMonkeyBusiness(monkeys))
}
