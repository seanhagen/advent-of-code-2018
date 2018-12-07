package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
--- Day 5: Alchemical Reduction ---
You've managed to sneak in to the prototype suit manufacturing lab. The
Elves are making decent progress, but are still struggling with the suit's
size reduction capabilities.

While the very latest in 1518 alchemical technology might have solved their
problem eventually, you can do better. You scan the chemical composition of
the suit's material and discover that it is formed by extremely long polymers
(one of which is available as your puzzle input).

The polymer is formed by smaller units which, when triggered, react with each
other such that two adjacent units of the same type and opposite polarity are
destroyed. Units' types are represented by letters; units' polarity is represented by
capitalization. For instance, r and R are units with the same type but opposite
polarity, whereas r and s are entirely different types and do not react.

For example:

In aA, a and A react, leaving nothing behind.
In abBA, bB destroys itself, leaving aA. As above, this then destroys itself, leaving nothing.
In abAB, no two adjacent units are of the same type, and so nothing happens.
In aabAAB, even though aa and AA are of the same type, their polarities match, and so nothing happens.
Now, consider a larger example, dabAcCaCBAcCcaDA:

dabAcCaCBAcCcaDA  The first 'cC' is removed.
dabAaCBAcCcaDA    This creates 'Aa', which is removed.
dabCBAcCcaDA      Either 'cC' or 'Cc' are removed (the result is the same).
dabCBAcaDA        No further actions can be taken.
After all possible reactions, the resulting polymer contains 10 units.

How many units remain after fully reacting the polymer you scanned?
*/

func main() {
	f := lib.LoadInput("../input.txt")

	input, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("unable to read from input file: %v\n", err)
		os.Exit(1)
	}

	out := processChain(string(input))

	fmt.Printf("units left after processing: %v\n", len(out)-1)
	// the '-1' is for the '\n' character
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
