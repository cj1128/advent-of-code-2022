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
