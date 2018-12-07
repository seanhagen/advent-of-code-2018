package lib

import (
	"bufio"
	"fmt"
	"os"
)

func LoadInput(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("unable to open input: %v\n", err)
		os.Exit(1)
	}
	return f
}

func LoopOverLines(file *os.File, fn func(line []byte) error) error {
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	for ; err == nil; line, _, err = r.ReadLine() {
		x := fn(line)
		if x != nil {
			fmt.Printf("\n\ngot error: %v\n", x)
			os.Exit(1)
		}
	}
	return err
}
