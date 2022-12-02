package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadStdin() string {
	buf, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(fmt.Errorf("could not read stdin: %v", err))
	}

	return string(buf)
}
