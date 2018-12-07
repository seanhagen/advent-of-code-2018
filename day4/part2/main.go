package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
--- Part Two ---
Strategy 2: Of all guards, which guard is most frequently asleep on the same minute?

In the example above, Guard #99 spent minute 45 asleep more than any other guard or minute - three times in total. (In all other cases, any guard spent any minute asleep at most twice.)

What is the ID of the guard you chose multiplied by the minute you chose? (In the above example, the answer would be 99 * 45 = 4455.)
*/

const dateFmtStr = "2006-01-02 15:04"

const awake = -1
const asleep = 1

type pickedGuard struct {
	id     string
	count  int
	minute int
}

func main() {
	f := lib.LoadInput("../input.txt")

	found := map[int64][]string{}

	begin := time.Date(3000, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(1, 1, 1, 0, 0, 0, 0, time.Local)

	lib.LoopOverLines(f, func(line []byte) error {
		bits := strings.Fields(string(line))

		datestr := fmt.Sprintf(
			"%v %v",
			strings.Replace(bits[0], "[", "", -1),
			strings.Replace(bits[1], "]", "", -1),
		)

		t, err := time.Parse(dateFmtStr, datestr)
		if err != nil {
			fmt.Printf("unable to parse date string: %v\n", err)
			os.Exit(1)
		}

		if t.Before(begin) {
			begin = t
		}

		if t.After(end) {
			end = t
		}
		found[t.Unix()] = bits[2:]
		return nil
	})

	test := map[string]bool{}
	asleepCount := map[string][]int{}

	var guardID string

	for x := begin; x.Unix() <= end.Unix(); x = x.Add(time.Minute) {
		t, ok := found[x.Unix()]
		if ok {
			date := x
			if date.Hour() != 0 {
				diff, _ := time.ParseDuration(fmt.Sprintf("%vm", 60-date.Minute()))
				date = date.Add(diff)
			}

			switch t[0] {
			case "Guard":
				guardID = strings.Replace(t[1], "#", "", -1)

				y, ok := test[guardID]
				if !ok {
					y = true
				}
				test[guardID] = y

				z, ok := asleepCount[guardID]
				if !ok {
					z = genShift()
					asleepCount[guardID] = fill(z, 0, 0)
				}

			case "wakes":
				a := asleepCount[guardID]
				asleepCount[guardID] = add(a, date.Minute(), -1)
			case "falls":
				a := asleepCount[guardID]
				asleepCount[guardID] = add(a, date.Minute(), +1)
			}
		}
	}

	highest := pickedGuard{"", 0, 0}
	for gid := range test {
		count := asleepCount[gid]

		for i, x := range count {
			if x > highest.count {
				highest.minute = i
				highest.count = x
				highest.id = gid
			}
		}
	}
	spew.Dump(highest)
}

func genShift() []int {
	len := 60
	out := make([]int, len)
	for i := range out {
		out[i] = awake
	}
	return out
}

func fill(in []int, start, fill int) []int {
	for i := start; i < len(in); i++ {
		in[i] = fill
	}
	return in
}

func add(in []int, start, add int) []int {
	for i := start; i < len(in); i++ {
		in[i] += add
	}
	return in
}
