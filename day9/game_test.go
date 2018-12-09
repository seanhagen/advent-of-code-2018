package day9

import (
	"fmt"
	"testing"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

func TestScoreExample(t *testing.T) {
	expect := 32
	o := RunGame(9, 25)

	if expect != o {
		t.Errorf("got wrong score for game; expect %v, got %v", expect, o)
	}
}

func TestScoreAgainstFile(t *testing.T) {
	f := lib.LoadInput("./fake.txt")

	scan := fmt.Sprintf("%v: high score is %%d", line)

	lib.LoopOverLines(f, func(in []byte) error {
		line := string(in)

		var numPlay, lastMarble, score int
		fmt.Sscanf(line, scan, &numPlay, &lastMarble, &score)
		o := RunGame(numPlay, lastMarble)

		if o != score {
			t.Errorf("got wrong score for game (players: %v, marbles: %v); expected %v, got %v", numPlay, lastMarble, score, o)
		}

		return nil
	})
}
