package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

/*
--- Day 4: Repose Record ---
You've sneaked into another supply closet - this time, it's across from
the prototype suit manufacturing lab. You need to sneak inside and fix
the issues with the suit, but there's a guard stationed outside the lab,
so this is as close as you can safely get.

As you search the closet for anything that might help, you discover that
you're not the first person to want to sneak in. Covering the walls, someone
has spent an hour starting every midnight for the past few months secretly
observing this guard post! They've been writing down the ID of the one guard
on duty that night - the Elves seem to have decided that one guard was enough
for the overnight shift - as well as when they fall asleep or wake up while
at their post (your puzzle input).

For example, consider the following records, which have already been organized into chronological order:

[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up

Timestamps are written using year-month-day hour:minute format. The guard falling
asleep or waking up is always the one whose shift most recently started. Because all
asleep/awake times are during the midnight hour (00:00 - 00:59), only the minute
portion (00 - 59) is relevant for those events.

Visually, these records show that the guards are asleep at these times:

Date   ID   Minute
            000000000011111111112222222222333333333344444444445555555555
            012345678901234567890123456789012345678901234567890123456789
11-01  #10  .....####################.....#########################.....
11-02  #99  ........................................##########..........
11-03  #10  ........................#####...............................
11-04  #99  ....................................##########..............
11-05  #99  .............................................##########.....

The columns are Date, which shows the month-day portion of the relevant day; ID,
which shows the guard on duty that day; and Minute, which shows the minutes during
which the guard was asleep within the midnight hour. (The Minute column's header
shows the minute's ten's digit in the first row and the one's digit in the second
row.) Awake is shown as ., and asleep is shown as #.

Note that guards count as asleep on the minute they fall asleep, and they count as
awake on the minute they wake up. For example, because Guard #10 wakes up at 00:25
on 1518-11-01, minute 25 is marked as awake.

If you can figure out the guard most likely to be asleep at a specific time, you
might be able to trick that guard into working tonight so you can have the best
chance of sneaking in. You have two strategies for choosing the best guard/minute
combination.

Strategy 1: Find the guard that has the most minutes asleep. What minute does that
guard spend asleep the most?

In the example above, Guard #10 spent the most minutes asleep, a total of 50
minutes (20+25+5), while Guard #99 only slept for a total of 30 minutes (10+10+10).
Guard #10 was asleep most during minute 24 (on two days, whereas any other minute
the guard was asleep was only seen on one day).

While this example listed the entries in chronological order, your entries are in the
order you found them. You'll need to organize them before they can be analyzed.

What is the ID of the guard you chose multiplied by the minute you chose? (In the
above example, the answer would be 10 * 24 = 240.)
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

	mostAsleep := 0

	for gid := range test {
		count := asleepCount[gid]

		total := 0
		for _, x := range count {
			total += x
		}

		if total > mostAsleep {
			mostAsleep = total

			highest = pickedGuard{gid, 0, 0}

			for i, x := range count {
				if x > highest.count {
					highest.minute = i
					highest.count = x
					highest.id = gid
				}
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

func mapIntToString(in []int) []string {
	out := []string{}
	for _, i := range in {
		x := strconv.Itoa(i)
		out = append(out, x)
	}
	return out
}
