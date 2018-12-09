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
	players := []player{}
	for i := 0; i < numPlayer; i++ {
		p := player{marbles: []int{}}
		players = append(players, p)
	}

	curPlayer := 2

	l := list.New()
	e := l.PushFront(0)
	e = l.InsertAfter(1, e)

	for i := 2; i < lastMarble; i++ {
		if i%23 == 0 {
			p := players[curPlayer-1]
			p.marbles = append(p.marbles, i)

			e := backSeven(l, e)
			p.marbles = append(p.marbles, e.Value.(int))

			o := e.Next()
			if o == nil {
				o = l.Front()
			}
			l.Remove(e)
			e = o

			players[curPlayer-1] = p
		} else {
			n := e.Next()
			if n == nil {
				n = l.Front()
			}

			e = l.InsertAfter(i, n)
		}

		curPlayer++
		if curPlayer > numPlayer {
			curPlayer = 1
		}
	}

	scores := map[int]int{}
	highest := 0

	for id, p := range players {
		s := sum(p.marbles)
		if s > highest {
			highest = s
		}
		scores[id] = s
	}

	return highest
}

func backSeven(l *list.List, e *list.Element) *list.Element {
	for i := 0; i < 7; i++ {
		e = e.Prev()
		if e == nil {
			e = l.Back()
		}
	}
	return e
}

func sum(in []int) int {
	s := 0
	if len(in) == 0 {
		return s
	}
	for _, v := range in {
		s += v
	}
	return s
}
