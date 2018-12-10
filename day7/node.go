package day7

import "fmt"

// Node point in graph
type Node struct {
	Name     string
	Requires []*Node

	start   int
	end     int
	working bool

	time int
}

// AddRequirement ...
func (n *Node) AddRequirement(m *Node) {
	n.Requires = append(n.Requires, m)
}

// Setup ...
func (n *Node) Setup(base int) {
	n.time = base
}

// Start ...
func (n *Node) Start(s, base int) {
	for i := range letters {
		l := string(letters[i])
		if l == n.Name {
			n.time = i + 1
		}
	}

	n.start = s
	n.end = s + n.time + base
	n.working = true

	fmt.Printf("node %v start %v end %v\n", n.Name, n.start, n.end)
}

// Done ...
func (n *Node) Done(s int) bool {
	if n == nil {
		return true
	}
	if s == n.end {
		n.working = false
		return true
	}
	return false
}

// MeetsRequirements ...
func (n *Node) MeetsRequirements(r []*Node) bool {
	allMet := make([]bool, len(n.Requires))

	for i, n := range n.Requires {
		for _, m := range r {
			if n == m {
				allMet[i] = true
			}
		}
	}

	met := true
	for _, t := range allMet {
		met = met && t
	}
	return met
}

// Test ...
func (n *Node) Test() []string {
	out := []string{}
	for _, x := range n.Requires {
		out = append(out, x.Name)
	}
	return out
}
