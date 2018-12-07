package day6

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

const marks = "ABCDEFGHIJKLMNOP1234567890!@#$%^&*()[]{}|;:',<.>/?`~"

const empty = -1 * 100 * 100 * 100
const doubleClaim = -1 * 1000 * 1000 * 1000

const closest = empty * 5

type row []int

type Board struct {
	width  int
	height int
	points []*point
	board  []row
}

// setup ...
func (b *Board) Setup(in *os.File) error {
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
	b.width++

	b.fill()
	return nil
}

// fill ...
func (b *Board) fill() {
	for y := 0; y <= b.height; y++ {
		r := row{}
		for x := 0; x <= b.width; x++ {
			r = append(r, empty)
		}
		b.board = append(b.board, r)
	}

	for _, p := range b.points {
		b.board[p.y][p.x] = p.id
	}
}

// print ...
func (b *Board) Print() string {
	out := ""
	for _, row := range b.board {
		for _, col := range row {
			switch col {
			case empty:
				out = fmt.Sprintf("%v%v", out, ".")
			case doubleClaim:
				out = fmt.Sprintf("%v%v", out, ".")
			case closest:
				out = fmt.Sprintf("%v%v", out, "#")
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
func (b *Board) Compute() {
	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			lowest := b.width * b.height
			found := map[int][]*point{}

			for _, p := range b.points {
				d := p.manhatdist(x, y)
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
			square := b.board[y][x]

			if len(bits) > 1 {
				square = doubleClaim
			} else {
				c := bits[0]
				if x == c.x && y == c.y {
					square = bits[0].id
				} else {
					square = bits[0].id * -1
				}

			}
			b.board[y][x] = square
		}
	}
}

// computePart2 ...
func (b *Board) ComputePart2(target int) {
	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			distances := []int{}

			for _, p := range b.points {
				d := p.manhatdist(x, y)
				distances = append(distances, d)
			}

			total := sumSlice(distances)
			if total < target {
				square := b.board[y][x]

				if square == empty {
					square = closest
				} else {
					for i, p := range b.points {
						if p.x == x && p.y == y {
							p.covered = true
						}
						b.points[i] = p
					}
				}

				b.board[y][x] = square
			}
		}
	}
}

// Part2Area ...
func (b *Board) Part2Area() int {
	out := 0

	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			square := b.board[y][x]
			if square == closest {
				out++
			}
		}
	}

	for _, p := range b.points {
		if p.covered {
			out++
		}
	}

	return out
}

// infinite ...
func (b Board) Infinite() []int {
	found := map[int]bool{}

	for y := 0; y <= b.height; y++ {
		switch y {
		case 0:
			fallthrough
		case b.height:
			for x := 0; x < b.width; x++ {
				id := int(math.Abs(float64(b.board[y][x])))
				if shouldAdd(id) {
					found[id] = true
				}
			}
		default:
			id := int(math.Abs(float64(b.board[y][0])))
			if shouldAdd(id) {
				found[id] = true
			}

			id = int(math.Abs(float64(b.board[y][b.width])))
			if shouldAdd(id) {
				found[id] = true
			}
		}
	}

	out := []int{}
	for id := range found {
		out = append(out, id)
	}

	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// finite ...
func (b Board) Finite() map[int]int {
	tmp := b.Infinite()
	out := map[int]int{}

	for _, p := range b.points {
		if !intIn(tmp, p.id) {
			out[p.id] = 0
		}
	}

	for y := 0; y <= b.height; y++ {
		for x := 0; x <= b.width; x++ {
			k := b.board[y]
			l := k[x]

			id := int(math.Abs(float64(l)))
			_, ok := out[id]
			if ok {
				out[id]++
			}

		}
	}

	return out
}

// largestFinite ...
func (b Board) LargestFinite() (int, int) {
	highest := 0
	foundID := 0

	tmp := b.Finite()
	for id, area := range tmp {
		if area > highest {
			highest = area
			foundID = id
		}
	}

	return foundID, highest
}

func shouldAdd(in int) bool {
	a := int(math.Abs(float64(empty)))
	b := int(math.Abs(float64(doubleClaim)))
	return a != in && b != in
}

func intIn(c []int, i int) bool {
	for _, x := range c {
		if x == i {
			return true
		}
	}
	return false
}

func sumSlice(in []int) int {
	out := 0
	for _, x := range in {
		out += x
	}
	return out
}
