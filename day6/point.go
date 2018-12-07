package day6

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	id      int
	x       int
	y       int
	covered bool
}

// parse ...
func (p *point) parse(in string) error {
	bits := strings.Split(in, ",")
	x, err := strconv.Atoi(strings.Replace(bits[0], " ", "", -1))
	if err != nil {
		return err
	}
	p.x = x

	y, err := strconv.Atoi(strings.Replace(bits[1], " ", "", -1))
	if err != nil {
		return err
	}
	p.y = y

	return nil
}

// manhatdist ...
func (p point) manhatdist(x, y int) int {
	a := float64(p.x - x)
	b := float64(p.y - y)
	return int(math.Abs(a) + math.Abs(b))
}
