package day7

import "fmt"

type worker struct {
	id  int
	sec int

	done bool

	node *Node
	log  []string
}

// setup ...
func (w *worker) setup(i int) {
	w.log = []string{}
	w.id = i
}

// done ...
func (w *worker) work(sec int) bool {
	w.sec = sec
	w.log = append(w.log, ".")

	if w.node == nil {
		return false
	}
	d := w.node.Done(sec)

	fmt.Printf("worker %v, node: %v, second: %v, node done: %v\n", w.id, w.node.Name, w.sec, d)

	if d {
		w.done = true
	} else {
		w.log[sec] = w.node.Name
	}

	return w.done
}

// setWorker ...
func (w *worker) setWorker(n *Node, s, base int) {
	n.Start(s, base)
	fmt.Printf("setting worker %v node to '%v', start: %v, end: %v\n", w.id, n.Name, s, n.end)
	w.log[w.sec] = n.Name
	// w.log = append(w.log, n.Name)
	w.node = n
	w.done = false
}

// needsWork ...
func (w worker) needsWork() bool {
	return w.done == true || w.node == nil
}
