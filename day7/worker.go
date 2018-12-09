package day7

import "fmt"

type worker struct {
	id  int
	sec int

	log []string

	node *Node
	prev *Node
}

// setup ...
func (w *worker) setup(i int) {
	w.log = []string{"."}
	w.id = i
}

// done ...
func (w *worker) work(sec int) {
	if w.node == nil {
		w.log = append(w.log, ".")
		fmt.Printf("worker %v, second %v, no node to work on\n", w.id, sec)
		return
	}
	w.log = append(w.log, w.node.Name)

	d := w.node.Done(sec)
	fmt.Printf("worker %v, second %v, node: %v, done: %v\n", w.id, sec, w.node.Name, d)
	if d {
		w.prev = w.node
		w.node = nil
		return
	}

	fmt.Printf("second: %v, len log: %v, log: %#v\n", sec, len(w.log), w.log)
}

// setWorker ...
func (w *worker) setWorker(n *Node, s, base int) {
	n.Start(s, base)
	w.node = n
	fmt.Printf("setting worker %v node to '%v', start: %v, end: %v, second: %v, log: %#v\n", w.id, n.Name, s, n.end, s, w.log)

	w.log[w.sec] = n.Name
}

// prevNode ...
func (w worker) prevNode() *Node {
	return w.prev
}

// needsWork ...
func (w worker) needsWork() bool {
	return w.node == nil
}
