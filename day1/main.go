package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("unable to open frequency file! reason: %v\n", err)
		os.Exit(1)
	}

	freq := 0

	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	for ; err == nil; line, _, err = r.ReadLine() {
		fmt.Printf("sign is '%v', value is %v\n", string(line[0]), string(line[1:]))
		sign := string(line[0])
		tmp := string(line[1:])
		val, err := strconv.Atoi(tmp)
		if err != nil {
			fmt.Printf("unable to parse frequency change: %v\n", err)
			os.Exit(1)
		}

		switch sign {
		case "-":
			freq -= val
		case "+":
			freq += val
		}
	}

	fmt.Printf("final frequency value: %v\n", freq)
}
