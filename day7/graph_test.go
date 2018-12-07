package day7

import (
	"strings"
	"testing"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

func TestFirstNode(t *testing.T) {
	f := lib.LoadInput("./fake.txt")
	g := &Graph{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}

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
	f := lib.LoadInput("./fake.txt")
	g := &Graph{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}

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
	f := lib.LoadInput("./fake.txt")
	g := &Graph{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}

	expect := "CABDFE"
	got := g.Print()

	if expect != got {
		t.Errorf("wrong output, expected '%v', got '%v'", expect, got)
	}
}

func TestMultipleStartNodes(t *testing.T) {
	f := lib.LoadInput("./fake2.txt")
	g := &Graph{}
	err := g.Setup(f)
	if err != nil {
		t.Errorf("got error setting up graph: %v", err)
		t.FailNow()
	}

	expect := "ABCED"
	got := g.Print()

	if expect != got {
		t.Errorf("wrong output, expected '%v', got '%v'", expect, got)
	}
}
