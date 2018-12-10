package day7

import (
	"fmt"
)

// SetupWork ...
func (g *Graph) SetupWork(workers, base int) {
	g.numWorkers = workers
	g.baseWorkTime = base
	start := g.FirstNodes()
	g.needsWork = append(g.needsWork, start...)

	g.workers = map[int]*worker{}
	for i := 1; i <= g.numWorkers; i++ {
		g.workers[i] = &worker{id: i}
		g.workers[i].setup(i)
	}
}

// DoWork ...
func (g *Graph) DoWork() int {
	i := 0

	fmt.Printf("base working time: %v\n", g.baseWorkTime)

	for ; g.workLeft(); i++ {
		for _, l := range g.locked {
			if l.MeetsRequirements(g.done) {
				g.needsWork = append(g.needsWork, l)
			}
		}
		g.needsWork = sortnodes(g.needsWork)

		for _, w := range g.workers {
			if w.node == nil {
				if len(g.needsWork) > 0 {
					x := g.needsWork[0]
					w.setWorker(x, i, g.baseWorkTime)
					g.needsWork = g.needsWork[1:]
				}
			}
			w.work(i, g)
		}

		if i >= 200 {
			break
		}

		if len(g.done) == len(g.nodes) {
			break
		}
	}

	// fmt.Printf("Second    Worker 1   Worker 2\n")
	// for j := 0; j < i; j++ {
	// 	fmt.Printf("%4v", j)
	// 	for _, w := range g.workers {
	// 		fmt.Printf("%10v", w.log[j])
	// 	}
	// 	fmt.Printf("\n")
	// }
	// fmt.Printf("\n\n")

	return i - 1
}

// WorkDone ...
func (g *Graph) workLeft() bool {
	return len(g.done) != len(g.nodes)
}

// PrintWork ...
func (g *Graph) PrintWork() string {
	return ""
}
