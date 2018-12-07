package day7

import (
	"testing"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

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
