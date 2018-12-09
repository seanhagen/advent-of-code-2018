package day8

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Tree ...
type Tree struct {
	n *Node
}

// Setup ...
func (t *Tree) Setup(in *os.File) error {
	tmp, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	o := strings.Replace(string(tmp), "\n", "", -1)
	bits := strings.Split(o, " ")

	vals := sliceToInt(bits)

	n := &Node{}
	_ = n.create(vals[0], vals[1], vals[2:])
	t.n = n

	return nil
}

// SumMeta ...
func (t *Tree) SumMeta() int {
	vals := t.n.meta()

	out := 0
	for _, x := range vals {
		out += x
	}

	return out
}

// SumRoot ...
func (t *Tree) SumRoot() int {
	return t.n.value()
}

func sliceToInt(in []string) []int {
	out := []int{}

	for _, i := range in {
		x, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		out = append(out, x)
	}

	return out
}
