package day7

import (
	"strings"
	"testing"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

func TestFirstNode(t *testing.T) {
	g := setup("./fake.txt", t)

	n := g.FirstNodes()
	tmp := []string{}

	for _, x := range n {
		tmp = append(tmp, x.Name)
	}
	got := strings.Join(tmp, "")

	expect := "C"

	if n == nil {
		t.Errorf("Expected a node, got nil")
	} else {
		if got != expect {
			t.Errorf("Unexpected node, expected '%v', got '%v'", expect, got)
		}
	}
}

func TestFirstNodeChildren(t *testing.T) {
	g := setup("./fake.txt", t)

	x := g.FirstNodes()
	n := x[0]

	expect := "AF"

	children := g.Children(n)

	tmp := []string{}
	for _, x := range children {
		tmp = append(tmp, x.Name)
	}
	got := strings.Join(tmp, "")

	if expect != got {
		t.Errorf("didn't get proper children, expect '%v', got '%v'", expect, got)
	}
}

func TestOutput(t *testing.T) {
	g := setup("./fake.txt", t)

	expect := "CABDFE"
	got := g.Print()

	if expect != got {
		t.Errorf("wrong output, expected '%v', got '%v'", expect, got)
	}
}

func TestMultipleStartNodes(t *testing.T) {
	g := setup("./fake2.txt", t)

	expect := "ABCED"
	got := g.Print()

	if expect != got {
		t.Errorf("wrong output, expected '%v', got '%v'", expect, got)
	}
}

func TestPart1(t *testing.T) {
	g := setup("./input.txt", t)

	expect := "EPWCFXKISTZVJHDGNABLQYMORU"
	got := g.Print()

	if expect != got {
		t.Errorf("part 1 is broken!\n\texpected '%v'\n\tgot:   '%v'", expect, got)
	}
}

func TestPart2Example(t *testing.T) {
	expect := `Second   Worker 1   Worker 2   Done
   0        C          .
   1        C          .
   2        C          .
   3        A          F       C
   4        B          F       CA
   5        B          F       CA
   6        D          F       CAB
   7        D          F       CAB
   8        D          F       CAB
   9        D          .       CABF
  10        E          .       CABFD
  11        E          .       CABFD
  12        E          .       CABFD
  13        E          .       CABFD
  14        E          .       CABFD
  15        .          .       CABFDE
`

	g := setup("./fake.txt", t)
	g.DoWork()

	got := g.PrintWork()

	if expect != got {
		t.Errorf("unexpected output!\nexpected:\n%v\n\ngot:\n%v\n\n", expect, got)
	}
}

func setup(path string, t *testing.T) *Graph {
	f := lib.LoadInput(path)
	g := &Graph{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}
	return g
}
