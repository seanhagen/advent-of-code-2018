package day7

import "fmt"

type worker struct {
	id  int
	sec int

	workDone int

	log []string

	node *Node

	children []*Node
}

// setup ...
func (w *worker) setup(i int) {
	w.log = []string{}
	w.id = i
}

// done ...
func (w *worker) work(sec int, g *Graph) {
	if w.node != nil && w.node.Done(sec) {
		g.done = append(g.done, w.node)
		children := g.Children(w.node)
		for _, c := range children {
			if c.MeetsRequirements(g.done) {
				in := false
				for _, x := range g.needsWork {
					if x == c {
						in = true
					}
				}
				if !in {
					g.needsWork = append(g.needsWork, c)
				}
			} else {
				in := false
				for _, x := range g.locked {
					if x == c {
						in = true
					}
				}
				if !in {
					g.locked = append(g.locked, c)
				}
			}
		}

		w.node = nil
		if len(g.needsWork) > 0 {
			n := g.needsWork[0]
			g.needsWork = g.needsWork[1:]
			n.Start(sec, g.baseWorkTime)
			w.node = n
		}
	}

	if w.node != nil {
		w.log = append(w.log, w.node.Name)
	} else {
		w.log = append(w.log, ".")
	}

}

// setWorker ...
func (w *worker) setWorker(n *Node, s, base int) {
	n.Start(s, base)
	w.node = n
	fmt.Printf("\ts:%v, setting worker %v node to '%v', start: %v, end: %v, second: %v, log: %#v\n", s, w.id, n.Name, s, n.end, s, w.log)

	w.workDone = n.end
	// w.log[w.sec] = n.Name
}
