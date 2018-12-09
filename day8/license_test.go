package day8

import (
	"testing"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
c m # # #  #  #  # # # # #  # m m m
2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
A----------------------------------
    B----------- C-----------
                     D-----


A setup(2, 3, []int{0 3 10 11 12 1 1 0 1 99 2 1 1 2})

  B setup(0, 3, []int{10 11 12 1 1 0 1 99 2 1 1 2})
    B meta: 10  11  12

    return []int{1 1 0 1 99 2 1 1 2}

  C setup(1, 1, []int{0 1 99 2 1 1 2})

    D setup(0, 1, []{99, 2, 1, 1, 2})
      D meta: 99

      return []int{2, 1, 1, 2}

    C meta: []int{2}

  A meta: []int{1, 1, 2}


*/

func TestMetadataSum(t *testing.T) {
	x := setup("./fake.txt", t)

	expect := 138
	got := x.SumMeta()

	if expect != got {
		t.Errorf("wrong sum, expected %v, got %v", expect, got)
	}
}

func TestRootValue(t *testing.T) {
	x := setup("./fake.txt", t)

	expect := 66
	got := x.SumRoot()

	if expect != got {
		t.Errorf("wrong value, expected %v, got %v", expect, got)
	}
}

func setup(path string, t *testing.T) *Tree {
	f := lib.LoadInput(path)
	g := &Tree{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}
	return g
}
