package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
--- Day 3: No Matter How You Slice It ---
The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit
(thanks to someone who helpfully wrote its box IDs on the wall of the warehouse in
the middle of the night). Unfortunately, anomalies are still affecting them - nobody
can even agree on how to cut the fabric.

The whole piece of fabric they're working on is a very large square - at least 1000
inches on each side.

Each Elf has made a claim about which area of fabric would be ideal for Santa's suit.
All claims have an ID and consist of a single rectangle with edges parallel to the
edges of the fabric. Each claim's rectangle is defined as follows:

The number of inches between the left edge of the fabric and the left edge of the rectangle.
The number of inches between the top edge of the fabric and the top edge of the rectangle.
The width of the rectangle in inches.
The height of the rectangle in inches.
A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches
from the left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall.
Visually, it claims the square inches of fabric represented by # (and ignores the
square inches of fabric represented by .) in the diagram below:

...........
...........
...#####...
...#####...
...#####...
...#####...
...........
...........
...........

The problem is that many of the claims overlap, causing two or more claims to cover part
of the same areas. For example, consider the following claims:

#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
Visually, these claim the following areas:

........
...2222.
...2222.
.11XX22.
.11XX22.
.111133.
.111133.
........

The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while
adjacent to the others, does not overlap either of them.)

If the Elves all proceed with their own plans, none of them will have enough fabric.
How many square inches of fabric are within two or more claims?
*/

// const file = "../fake.txt"
const file = "../input.txt"

// const expect = `........
// ...2222.
// ...2222.
// .11XX22.
// .11XX22.
// .111133.
// .111133.
// ........
// `

const boardWidth = 1500
const boardHeight = 1500

// const boardWidth = 8
// const boardHeight = 7

const xspot = -2

type row map[int]int

type board struct {
	b      map[int]row
	claims map[int]*claim
}

// setup ...
func (b *board) setup() {
	b.claims = map[int]*claim{}
	b.b = map[int]row{}

	for i := 0; i <= boardHeight; i++ {
		b.b[i] = row{}
		for j := 0; j < boardWidth; j++ {
			b.b[i][j] = -1
		}
	}
}

// add ...
func (b board) add(c *claim) {
	b.claims[c.id] = c

	for i := c.y; i < c.y+c.h; i++ {
		for j := c.x; j < c.x+c.w; j++ {
			t := b.b[i][j]
			switch t {
			case -1:
				t = c.id
			default:
				t = xspot
			}
			b.b[i][j] = t
		}
	}
}

// print ...
func (b board) out() string {
	out := ""
	for i := 0; i <= boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			t := b.b[i][j]
			switch b.b[i][j] {
			case -1:
				out = fmt.Sprintf("%v%v", out, ".")
			case xspot:
				out = fmt.Sprintf("%v%v", out, "X")
			default:
				out = fmt.Sprintf("%v%v", out, t)
			}
		}
		out = fmt.Sprintf("%v%v", out, "\n")
	}
	return out
}

// count ...
func (b board) count() int {
	overlap := 0

	for i := 0; i <= boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			t := b.b[i][j]
			if t == xspot {
				overlap++
			}
		}
	}

	return overlap
}

type claim struct {
	x  int
	y  int
	w  int
	h  int
	id int
}

// parse ...
func (c *claim) parse(in string) error {
	bits := strings.Fields(in)
	id := bits[0]
	pos := bits[2]
	size := bits[3]

	id = strings.Replace(id, "#", "", -1)
	pos = strings.Replace(pos, ":", "", -1)

	bits = strings.Split(pos, ",")
	x := bits[0]
	y := bits[1]

	bits = strings.Split(size, "x")
	w := bits[0]
	h := bits[1]

	o, err := strconv.Atoi(x)
	if err != nil {
		return err
	}
	c.x = o

	o, err = strconv.Atoi(y)
	if err != nil {
		return err
	}
	c.y = o

	o, err = strconv.Atoi(w)
	if err != nil {
		return err
	}
	c.w = o

	o, err = strconv.Atoi(h)
	if err != nil {
		return err
	}
	c.h = o

	o, err = strconv.Atoi(id)
	if err != nil {
		return err
	}
	c.id = o

	return nil
}

func main() {
	f := lib.LoadInput("../input.txt")

	b := &board{}
	b.setup()

	lib.LoopOverLines(f, func(line []byte) error {
		c := &claim{}
		c.parse(string(line))
		b.add(c)
		return nil
	})

	fmt.Printf("overlapping inches: %v\n\n", b.count())
}
