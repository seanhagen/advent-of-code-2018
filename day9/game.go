package day9

import (
	"container/list"
)

const line = "%d players; last marble is worth %d points"

type player struct {
	marbles []int
}

// RunGame ...
func RunGame(numPlayer, lastMarble int) int {
	players := make([]int, numPlayer)

	l := list.New()
	e := l.PushFront(0)

	for i := 1; i < lastMarble; i++ {
		if i%23 == 0 {
			rm := e
			for j := 0; j < 7; j++ {
				rm = rm.Prev()
				if rm == nil {
					rm = l.Back()
				}
			}

			curPlayer := i % numPlayer
			players[curPlayer] += i + rm.Value.(int)
			e = rm.Next()
			l.Remove(rm)
		} else {
			n := e.Next()
			if n == nil {
				n = l.Front()
			}

			e = l.InsertAfter(i, n)
		}
	}

	highest := 0
	for _, p := range players {
		if p > highest {
			highest = p
		}
	}

	return highest
}
