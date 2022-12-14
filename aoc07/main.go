package main

import (
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"

	"cjting.me/aoc2022/utils"
)

func main() {
	stdin := utils.ReadStdin()
	fs := parse(stdin)

	// 1297683
	// part1(fs)

	// 5756764
	part2(fs)
}

// if children is nil, then it's a file, otherwise it's a directory
type FileItem struct {
	path     string
	name     string
	size     int
	children []*FileItem
}

func parse(str string) *FileItem {
	commands := parseCommands(str)

	cwd := ""
	m := make(map[string]*FileItem)

	for _, cmd := range commands {
		switch cmd.name {
		case "cd":
			{
				switch cmd.arg {
				case "/":
					{
						cwd = "/"
					}
				default:
					{
						cwd = path.Join(cwd, cmd.arg)
					}
				}
			}

		case "ls":
			{
				cur := m[cwd]
				if cur == nil {
					cur = &FileItem{path: cwd, name: path.Base(cwd)}
					m[cwd] = cur
				}

				for _, l := range cmd.output {
					if strings.HasPrefix(l, "dir ") {
						name := l[4:]
						sub := path.Join(cwd, name)
						item := &FileItem{path: sub, name: name}
						cur.children = append(cur.children, item)
						m[sub] = item
					} else {
						parts := strings.Split(l, " ")

						if len(parts) != 2 {
							utils.Err("invalid output %s", l)
						}

						size, err := strconv.Atoi(parts[0])
						if err != nil {
							utils.Err("invalid size %s", l)
						}

						name := parts[1]

						cur.children = append(cur.children, &FileItem{
							path: path.Join(cwd, name),
							name: name,
							size: size,
						})
					}
				}
			}
		}
	}

	root := m["/"]

	// calculate directory size
	var recur func(fs *FileItem)
	recur = func(fs *FileItem) {
		if fs.children == nil {
			return
		}

		for _, c := range fs.children {
			recur(c)
			fs.size += c.size
		}
	}

	recur(root)

	return root
}

type Command struct {
	// cd, ls
	name   string
	arg    string
	output []string
}

type CommandReader struct {
	lines []string
	idx   int
}

func (cr *CommandReader) readCommand() *Command {
	line, ok := cr.readLine()

	if !ok {
		return nil
	}

	if !isCommand(line) {
		utils.Err("invalid command %v", line)
	}

	result := &Command{}

	parts := strings.Split(line[2:], " ")

	result.name = parts[0]

	if result.name == "cd" {
		result.arg = parts[1]
	}

	if result.name == "ls" {
		result.output = cr.readTillNextCommand()
	}

	return result
}

func (cr *CommandReader) readLine() (string, bool) {
	if cr.idx >= len(cr.lines) {
		return "", false
	}

	result := cr.lines[cr.idx]
	cr.idx += 1
	return result, true
}

func (cr *CommandReader) readTillNextCommand() []string {
	var result []string

	for {
		line, ok := cr.readLine()

		if !ok {
			break
		}

		if isCommand(line) {
			cr.idx -= 1
			break
		}

		result = append(result, line)
	}

	return result
}

func newCommandReader(str string) CommandReader {
	lines := strings.Split(strings.TrimSpace(str), "\n")
	return CommandReader{lines, 0}
}

func parseCommands(str string) []*Command {
	var result []*Command

	cr := newCommandReader(str)

	for {
		cmd := cr.readCommand()
		if cmd == nil {
			break
		}

		result = append(result, cmd)
	}

	return result
}

func printFS(fs *FileItem) {
	_printFS(fs, "")
}

func _printFS(fs *FileItem, indent string) {
	fmt.Printf("%s%s %d\n", indent, fs.name, fs.size)
	for _, c := range fs.children {
		_printFS(c, indent+"  ")
	}
}

func part1(fs *FileItem) {
	// printFS((fs))

	result := 0

	var visit func(fs *FileItem)

	visit = func(fs *FileItem) {
		if fs.children == nil {
			return
		}

		if fs.size <= 100000 {
			result += fs.size
		}

		for _, c := range fs.children {
			visit(c)
		}
	}

	visit(fs)

	fmt.Println(result)
}

func isCommand(str string) bool {
	return strings.HasPrefix(str, "$ ")
}

func part2(fs *FileItem) {
	//
	needed := 30000000 - (70000000 - fs.size)

	if needed <= 0 {
		fmt.Println("no need to free any space")
	}

	var dirSizes []int

	var visit func(fs *FileItem)

	visit = func(fs *FileItem) {
		if fs.children == nil {
			return
		}

		dirSizes = append(dirSizes, fs.size)

		for _, c := range fs.children {
			visit(c)
		}
	}

	visit(fs)

	sort.Slice(dirSizes, func(i, j int) bool {
		return dirSizes[i] < dirSizes[j]
	})

	for _, s := range dirSizes {
		if s >= needed {
			fmt.Println(s)
			break
		}
	}
}
