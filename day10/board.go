package day10

import (
	"fmt"
	"io"
	"math"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
#...#..###
#...#...#.
#...#...#.
#####...#.
#...#...#.
#...#...#.
#...#...#.
#...#..###
*/

var H = [][]int{
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 1, 1, 1, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 0, 0, 1},
}

var I = [][]int{
	[]int{1, 0, 1},
	[]int{0, 1, 0},
	[]int{0, 1, 0},
	[]int{0, 1, 0},
	[]int{0, 1, 0},
	[]int{0, 1, 0},
	[]int{0, 1, 0},
	[]int{1, 1, 1},
}

// Board ...
type Board struct {
	points []*point

	minX int
	minY int
	maxX int
	maxY int
}

// Setup ...
func (b *Board) Setup(path string) error {
	f := lib.LoadInput(path)

	b.minX = 0
	b.minY = 0
	b.maxX = 0
	b.maxY = 0

	err := lib.LoopOverLines(f, func(line []byte) error {
		p := &point{}
		p.setup(string(line))

		if p.x > b.maxX {
			b.maxX = p.x
		}
		if p.y > b.maxY {
			b.maxY = p.y
		}
		if p.x < b.minX {
			b.minX = p.x
		}
		if p.y < b.minX {
			b.minY = p.y
		}

		b.points = append(b.points, p)
		// fmt.Printf("point: <%v, %v>\n", p.x, p.y)

		return nil
	})

	if err != io.EOF {
		return err
	}

	return nil
}

// Print ...
func (b Board) Print() string {
	out := ""
	for i := b.minY; i <= b.maxY; i++ {
		for j := b.minX; j <= b.maxX; j++ {
			z := "."
			for _, p := range b.points {
				if p.x == j && p.y == i {
					z = "#"
				}
			}
			out = fmt.Sprintf("%v%v", out, z)
		}
		out = fmt.Sprintf("%v%v", out, "\n")
	}

	return out
}

// Message ...
func (b Board) Message() string {
	for i := 0; ; i++ {
		switch i {
		case 5:
			fmt.Printf("A\n")
		case 6:
			fmt.Printf("AA\n")
		case 7:
			fmt.Printf("AAA\n")
		case 10:
			fmt.Printf("B\n")
		case 50:
			fmt.Printf("C\n")
		case 100:
			fmt.Printf("D\n")
		case 500:
			fmt.Printf("E\n")
		case 1000:
			fmt.Printf("F\n")
		default:
			fmt.Printf("%v\n", i)
		}

		for _, p := range b.points {
			foundH := b.TestForLetter(p, H)
			foundI := b.TestForLetter(p, I)

			if foundH || foundI {
				return b.Print()
			}
		}
		b.step()

		if i > 5 {
			return b.Print()
		}
	}
}

// TestForH ...
func (b Board) TestForLetter(p *point, test [][]int) bool {
	should := []point{}

	height := len(test)
	width := len(test[0])
	count := 0

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if test[i][j] == 1 {
				t := point{
					x: p.x + j,
					y: p.y + i,
				}
				should = append(should, t)

				x := p.x + j
				y := p.y + i

				for _, p := range b.points {
					if p.x == x && p.y == y {
						count++
					}
				}
			}
		}
	}

	return count == len(should)
}

// DoSteps ...
func (b *Board) DoSteps(i int) {
	for j := 1; j < i+1; j++ {
		b.step()
	}
}

// step ...
func (b *Board) step() {
	for i := 0; i < len(b.points); i++ {
		b.points[i].step()
	}
}

func addAbs(a, b int) int {
	x := int(math.Abs(float64(a)))
	y := int(math.Abs(float64(b)))
	return x + y
}
