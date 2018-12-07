package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/advent-of-code-2018/lib"
)

func TestBoard(t *testing.T) {
	expect := `..........
.A........
..........
........C.
...D......
.....E....
.B........
..........
..........
........F.`

	f := lib.LoadInput("../fake.txt")

	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("unable to read from input file: %v\n", err)
		os.Exit(1)
	}

	spew.Dump(input)
	spew.Dump(expect)
}
