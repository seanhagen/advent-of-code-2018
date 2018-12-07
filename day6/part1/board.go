package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

const marks = "ABCDEFGHIJKLMNOP1234567890!@#$%^&*()[]{}|;:',<.>/?`~`"

const empty = -1 * 100 * 100 * 100
const doubleClaim = -1 * 1000 * 1000 * 1000

type row []int

type board struct {
	width  int
	height int
	points []*point
	board  []row
}

// setup ...
func (b *board) setup(in *os.File) error {
	b.board = []row{}

	b.points = []*point{}

	id := 1

	err := lib.LoopOverLines(in, func(line []byte) error {
		p := &point{
			id: id,
		}
		id++

		err := p.parse(string(line))
		if err != nil {
			return nil
		}

		if p.x > b.width {
			b.width = p.x
		}

		if p.y > b.height {
			b.height = p.y
		}

		b.points = append(b.points, p)
		return nil
	})

	if err != nil && err != io.EOF {
		return err
	}

	b.fill()
	return nil
}

// fill ...
func (b *board) fill() {
	for i := 0; i <= b.height; i++ {
		r := row{}
		for j := 0; j <= b.width+1; j++ {
			r = append(r, empty)
		}
		b.board = append(b.board, r)
	}

	for _, p := range b.points {
		b.board[p.y][p.x] = p.id
	}
}

// print ...
func (b *board) print() string {
	out := ""
	for _, row := range b.board {
		for _, col := range row {
			switch col {
			case empty:
				out = fmt.Sprintf("%v%v", out, ".")
			case doubleClaim:
				out = fmt.Sprintf("%v%v", out, ".")
			default:
				idx := int(math.Abs(float64(col))) - 1
				x := string(marks[idx])
				if col < 0 {
					x = strings.ToLower(x)
				}
				out = fmt.Sprintf("%v%v", out, x)
			}
		}
		out = fmt.Sprintf("%v%v", out, "\n")
	}
	return out
}

// compute ...
func (b *board) compute() {
	for i := 0; i <= b.height; i++ {
		for j := 0; j <= b.width+1; j++ {
			lowest := b.width * b.height
			found := map[int][]*point{}

			for _, p := range b.points {
				d := p.manhatdist(j, i)
				if d < lowest {
					lowest = d
				}

				tmp, ok := found[d]
				if !ok {
					tmp = []*point{}
				}
				tmp = append(tmp, p)
				found[d] = tmp
			}

			bits := found[lowest]

			square := b.board[i][j]

			if len(bits) > 1 {
				square = doubleClaim
			} else {
				c := bits[0]
				if j == c.x && i == c.y {
					square = bits[0].id
				} else {
					square = bits[0].id * -1
				}

			}
			b.board[i][j] = square
		}
	}
}
