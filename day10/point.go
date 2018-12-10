package day10

import (
	"strconv"
	"strings"
)

const line = "position=<%d, %d> velocity=<%1d, %d>"

type point struct {
	x    int
	y    int
	xvel int
	yvel int

	check bool
}

// setup ...
func (p *point) setup(in string) {
	// fmt.Printf("got line: \n%v\n", in)
	in = strings.Replace(in, "<", "", -1)
	in = strings.Replace(in, ">", "", -1)
	in = strings.Replace(in, ",", "", -1)
	in = strings.Replace(in, "=", " ", -1)

	bits := strings.Split(in, " ")

	o := []int{}
	for _, x := range bits {
		switch x {
		case "position":
		case "velocity":
		case "":
		default:
			y, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			o = append(o, y)
		}
	}

	p.x = o[0]
	p.y = o[1]
	p.xvel = o[2]
	p.yvel = o[3]

	// spew.Dump(p)
	// os.Exit(1)
}

// step ...
func (p *point) step() {
	p.x += p.xvel
	p.y += p.yvel
}

// back ...
func (p *point) back() {
	p.x -= p.xvel
	p.y -= p.yvel
}
