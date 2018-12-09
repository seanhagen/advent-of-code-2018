package day8

// Node ...
type Node struct {
	children []*Node
	metadata []int
}

// create ...
func (n *Node) create(numChild, numMeta int, data []int) []int {
	n.children = make([]*Node, numChild)
	n.metadata = make([]int, numMeta)

	if numChild > 0 {
		for i := 0; i < numChild; i++ {
			nc := data[0]
			nm := data[1]
			data = data[2:]
			c := &Node{}
			data = c.create(nc, nm, data)
			n.children[i] = c
		}
	}

	if numMeta > 0 {
		for i := 0; i < numMeta; i++ {
			n.metadata[i] = data[i]
		}
		data = data[numMeta:]
	}

	return data
}

// meta ...
func (n *Node) meta() []int {
	vals := []int{}

	for _, c := range n.children {
		tmp := c.meta()
		vals = append(vals, tmp...)
	}

	vals = append(vals, n.metadata...)
	return vals
}

// value ...
func (n *Node) value() int {
	if len(n.children) == 0 {
		v := 0
		for _, m := range n.metadata {
			v += m
		}
		return v
	}

	v := 0
	for _, m := range n.metadata {
		if m == 0 {
			continue
		}

		if m-1 >= len(n.children) {
			continue
		}
		c := n.children[m-1]

		x := c.value()
		v += x
	}
	return v
}
