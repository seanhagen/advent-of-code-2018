package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
--- Part Two ---
Time to improve the polymer.

One of the unit types is causing problems; it's preventing the polymer from collapsing as
much as it should. Your goal is to figure out which unit type is causing the most problems,
remove all instances of it (regardless of polarity), fully react the remaining polymer,
and measure its length.

For example, again using the polymer dabAcCaCBAcCcaDA from above:

Removing all A/a units produces dbcCCBcCcD. Fully reacting this polymer produces dbCBcD, which has length 6.
Removing all B/b units produces daAcCaCAcCcaDA. Fully reacting this polymer produces daCAcaDA, which has length 8.
Removing all C/c units produces dabAaBAaDA. Fully reacting this polymer produces daDA, which has length 4.
Removing all D/d units produces abAcCaCBAcCcaA. Fully reacting this polymer produces abCBAc, which has length 6.

In this example, removing all C/c units was best, producing the answer 4.

What is the length of the shortest polymer you can produce by removing all units of exactly one type and fully reacting the result?
*/

func main() {
	f := lib.LoadInput("../input.txt")

	tmp, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("unable to read from input file: %v\n", err)
		os.Exit(1)
	}
	input := string(tmp[:len(tmp)-1])
	types := getTypes(input)

	lengths := map[string]int{}

	for _, t := range types {
		o := removePotentialProblem(input, t)
		lengths[t] = len(o)
	}

	shortest := len(input)
	remove := ""

	for r, l := range lengths {
		if l < shortest {
			shortest = l
			remove = r
		}
	}

	spew.Dump(shortest, remove)

}

func getTypes(in string) []string {
	bits := strings.Split(in, "")
	found := map[string]bool{}

	for _, b := range bits {
		x := strings.ToLower(b)
		found[x] = true
	}

	out := []string{}

	for k := range found {
		out = append(out, k)
	}

	return out
}

func removePotentialProblem(in, p string) string {
	in = strings.Replace(in, p, "", -1)
	in = strings.Replace(in, strings.ToUpper(p), "", -1)
	return processChain(in)
}

func processChain(in string) string {
begin:
	start := len(in)
	out := pass(in)

	if len(out) == 0 {
		return out
	}

	if len(out) == start {
		return out
	}

	in = out
	goto begin
}

func pass(in string) string {
	for i := range in {
		if i == 0 {
			continue
		}
		if i >= len(in) {
			break
		}
		if in[i]^in[i-1] == 32 {
			in = in[:i-1] + in[i+1:]
			i = -1
		}
	}
	return in
}
