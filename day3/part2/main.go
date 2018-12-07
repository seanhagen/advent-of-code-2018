package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
Amidst the chaos, you notice that exactly one claim doesn't overlap by even
a single square inch of fabric with any other claim. If you can somehow draw
attention to it, maybe the Elves will be able to make Santa's suit after all!

For example, in the claims above, only claim 3 is intact after all claims are made.

What is the ID of the only claim that doesn't overlap?
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

const blank = -1
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
			b.b[i][j] = blank
		}
	}
}

// notOverlapping ...
func (b board) notOverlapping() []int {
	out := []int{}
	for id, c := range b.claims {
		if !c.overlapping {
			out = append(out, id)
		}
	}
	return out
}

// add ...
func (b *board) add(c *claim) {
	for i := c.y; i < c.y+c.h; i++ {
		for j := c.x; j < c.x+c.w; j++ {
			t := b.b[i][j]
			switch t {
			case blank:
				t = c.id
			default:
				c.overlapping = true
				old, ok := b.claims[t]
				if ok {
					old.overlapping = true
					b.claims[old.id] = old
				}
				t = xspot
			}
			b.b[i][j] = t
		}
	}
	b.claims[c.id] = c
}

// print ...
func (b board) out() string {
	out := ""
	for i := 0; i <= boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			t := b.b[i][j]
			switch b.b[i][j] {
			case blank:
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
	x           int
	y           int
	w           int
	h           int
	id          int
	overlapping bool
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

	c.overlapping = false
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

	// out := b.out()
	// fmt.Printf("got: \n%v\n\n", out)
	// fmt.Printf("overlapping inches: %v\n\n", b.count())

	over := b.notOverlapping()
	fmt.Printf("claim ids for claims that don't overlap: %v\n\n", over)
}
