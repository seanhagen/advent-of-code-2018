package day6

import (
	"sort"
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

	f := lib.LoadInput("./fake.txt")
	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	got := b.Print()

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

	f := lib.LoadInput("./fake.txt")

	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	b.Compute()
	got := b.Print()

	// fmt.Printf("\n\n%v\n\n", got)

	if expect != got {
		t.Errorf("got bad board:\n expected:\n%v\n\ngot:\n%v\n\n", expect, got)
	}

}

func TestInfinite(t *testing.T) {
	should := []int{6, 1, 2, 3}

	f := lib.LoadInput("./fake.txt")
	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	b.Compute()
	inf := b.Infinite()
	got := []int{}
	for _, p := range inf {
		got = append(got, p)
	}

	if !arEq(should, got) {
		t.Errorf("inf not equal. \nshould have %v\ngot %v", should, inf)
	}
}

func TestFindLargest(t *testing.T) {
	f := lib.LoadInput("./fake.txt")
	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}
	b.Compute()

	should := 5
	area := 17

	foundID, highest := b.LargestFinite()

	if foundID != should {
		t.Errorf("didn't find largest safe area, expected %v got %v", should, foundID)
	}

	if highest != area {
		t.Errorf("found area not correct, expected %v got %v", area, highest)
	}
}

func TestSafe(t *testing.T) {
	expect := `..........
.A........
..........
...###..C.
..#D###...
..###E#...
.B.###....
..........
..........
........F.
`

	f := lib.LoadInput("./fake.txt")
	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	dist := 32

	b.ComputePart2(dist)

	got := b.Print()

	if got != expect {
		t.Errorf("didn't get right output!\n\nexpect:\n%v\n\ngot:\n%v\n\n", expect, got)
	}
}

func TestSafeArea(t *testing.T) {
	f := lib.LoadInput("./fake.txt")
	b := &Board{}
	err := b.Setup(f)
	if err != nil {
		t.Errorf("unable to setup board: %v", err)
	}

	dist := 32

	b.ComputePart2(dist)

	expect := 16

	got := b.Part2Area()

	if got != expect {
		t.Errorf("wrong area, expected %v got %v", expect, got)
	}
}

func arEq(a, b []int) bool {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
