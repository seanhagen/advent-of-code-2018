package main

import (
	"fmt"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

/*

--- Part Two ---
Confident that your list of box IDs is complete, you're ready to find the boxes
full of prototype fabric.

The boxes will have IDs which differ by exactly one character at the same position
in both strings. For example, given the following box IDs:

abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz

The IDs abcde and axcye are close, but they differ by two characters (the second and
fourth). However, the IDs fghij and fguij differ by exactly one character, the third
(h and u). Those must be the correct boxes.

What letters are common between the two correct box IDs? (In the example above, this
is found by removing the differing character from either ID, producing fgij.)

*/

func main() {
	ids := []string{}
	opts := levenshtein.Options{
		InsCost: 2,
		DelCost: 2,
		SubCost: 1,
		Matches: func(sourceCharacter rune, targetCharacter rune) bool {
			return sourceCharacter == targetCharacter
		},
	}

	f := lib.LoadInput("../input.txt")

	lib.LoopOverLines(f, func(line []byte) error {
		ids = append(ids, string(line))
		return nil
	})

	var a, b string

	for _, s := range ids {
		for _, t := range ids {
			if x := dist(s, t, opts); x == 1 {
				a = s
				b = t
				goto done
			}
		}
	}

done:
	fmt.Printf("two ids: \n\t%v\n\t%v\n\n", a, b)

	both := same(a, b)
	fmt.Printf("in both: \n\t%v\n\n", both)
}

func same(a, b string) string {
	aparts := strings.Split(a, "")
	bparts := strings.Split(b, "")

	s := ""

	for i, x := range aparts {
		if x == bparts[i] {
			s = fmt.Sprintf("%v%v", s, x)
		}
	}

	return s
}

func dist(a, b string, opt levenshtein.Options) int {
	if a == b {
		return 100
	}
	return levenshtein.DistanceForStrings([]rune(a), []rune(b), opt)
}
