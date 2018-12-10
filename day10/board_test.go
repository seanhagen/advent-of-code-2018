package day10

import (
	"testing"
)

func TestWidth(t *testing.T) {
	expect := 2
	got := addAbs(-1, 1)
	if expect != got {
		t.Errorf("nope. expect: %v, got: %v", expect, got)
	}
}

func TestInitial(t *testing.T) {
	expect := `........#.............
................#.....
.........#.#..#.......
......................
#..........#.#.......#
...............#......
....#.................
..#.#....#............
.......#..............
......#...............
...#...#.#...#........
....#..#..#.........#.
.......#..............
...........#..#.......
#...........#.........
...#.......#..........
`

	b := &Board{}
	b.Setup("./test.txt")

	o := b.Print()

	if o != expect {
		t.Errorf("wrong output!\nexpected: \n%v\n\ngot: %v\n", expect, o)
	}
}

func TestOneSecond(t *testing.T) {
	expect := `......................
......................
..........#....#......
........#.....#.......
..#.........#......#..
......................
......#...............
....##.........#......
......#.#.............
.....##.##..#.........
........#.#...........
........#...#.....#...
..#...........#.......
....#.....#.#.........
......................
......................
`

	b := &Board{}
	b.Setup("./test.txt")

	b.DoSteps(1)

	o := b.Print()

	if o != expect {
		t.Errorf("wrong output!\nexpected: \n%v\n\ngot: \n%v\n", expect, o)
	}
}

func TestHasH(t *testing.T) {
	expect := true

	b := &Board{}
	b.Setup("./test.txt")

	b.DoSteps(3)

	got := false

	for _, p := range b.points {
		x := b.TestForLetter(p, H)
		if x {
			got = true
			break
		}
	}

	if got != expect {
		t.Errorf("wrong message. expected '%v', got: '%v'", expect, got)
	}
}

func TestHasI(t *testing.T) {
	expect := true

	b := &Board{}
	b.Setup("./test.txt")

	b.DoSteps(3)

	got := false

	for _, p := range b.points {
		x := b.TestForLetter(p, I)
		if x {
			got = true
			break
		}
	}

	if got != expect {
		t.Errorf("wrong message. expected '%v', got: '%v'", expect, got)
	}
}

func TestFindSmallestBoudingBox(t *testing.T) {
	b := &Board{}
	b.Setup("./test.txt")

	area := b.FindSmallestBoundingBox()
	expect := 63
	if area != expect {
		t.Errorf("wrong area, expected %v got %v", expect, area)
	}
}

func TestSmallestMessage(t *testing.T) {
	b := &Board{}
	b.Setup("./test.txt")

	_ = b.FindSmallestBoundingBox()

	expect := `#...#..###
#...#...#.
#...#...#.
#####...#.
#...#...#.
#...#...#.
#...#...#.
#...#..###
`

	out := b.Print()
	if expect != out {
		t.Errorf("wrong message. \nexpected:\n%v\n\ngot: \n%v\n", expect, out)
	}
}

func TestFindMessage(t *testing.T) {
	expect := `......................
......................
......................
......................
......#...#..###......
......#...#...#.......
......#...#...#.......
......#####...#.......
......#...#...#.......
......#...#...#.......
......#...#...#.......
......#...#..###......
......................
......................
......................
......................
`

	b := &Board{}
	b.Setup("./test.txt")

	out := b.Message()

	if expect != out {
		t.Errorf("wrong message.\nexpect: \n%v\n\ngot:\n%v\n", expect, out)
	}

}
