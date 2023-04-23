package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadStdin() string {
	buf, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(fmt.Errorf("could not read stdin: %v", err))
	}

	return string(buf)
}

func Err(f string, args ...interface{}) {
	panic(fmt.Errorf(f, args...))
}

func MustAtoi(f string) int {
	num, err := strconv.Atoi(f)

	if err != nil {
		Err("invalid number string: %s", f)
	}

	return num
}

func MustAtoi64(f string) int64 {
	num, err := strconv.ParseInt(f, 10, 64)

	if err != nil {
		Err("invalid number string: %s", f)
	}

	return num
}

func FindCommonInString(strs []string) []rune {
	if len(strs) == 0 {
		return nil
	}

	ms := make([]map[rune]bool, len(strs))
	for i := 0; i < len(strs); i++ {
		ms[i] = make(map[rune]bool)
	}

	for idx, str := range strs {
		for _, rune := range str {
			ms[idx][rune] = true
		}
	}

	var result []rune

outer:
	for k := range ms[0] {
		for _, m := range ms[1:] {
			if !m[k] {
				continue outer
			}
		}

		result = append(result, k)
	}

	return result
}

func Assert(expr bool) {
	if !expr {
		panic("assert error")
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
