package main

import (
	"testing"

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
........F.
`

	f := lib.LoadInput("../fake.txt")
	b := &board{}
	err := b.setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	got := b.print()

	if expect != got {
		t.Errorf("got bad board:\n expected:\n%v\n\ngot:\n%v\n\n", expect, got)
	}
}

func TestCompute(t *testing.T) {
	expect := `aaaaa.cccc
aAaaa.cccc
aaaddecccc
aadddeccCc
..dDdeeccc
bb.deEeecc
bBb.eeee..
bbb.eeefff
bbb.eeffff
bbb.ffffFf
`

	f := lib.LoadInput("../fake.txt")

	b := &board{}
	err := b.setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	b.compute()

	got := b.print()

	// fmt.Printf("\n\n%v\n\n", got)

	if expect != got {
		t.Errorf("got bad board:\n expected:\n%v\n\ngot:\n%v\n\n", expect, got)
	}

}
