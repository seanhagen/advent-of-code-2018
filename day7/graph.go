package day7

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

// https://flaviocopes.com/golang-data-structure-graph/
// https://godoc.org/github.com/gonum/graph

// Node point in graph
type Node struct {
	Name     string
	Requires []*Node
}

// AddRequirement ...
func (n *Node) AddRequirement(m *Node) {
	n.Requires = append(n.Requires, m)
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

// Graph data structure
type Graph struct {
	nodes map[string]*Node
	edges map[*Node][]*Node
}

// Last ...
func (g *Graph) Last() *Node {
	for _, n := range g.nodes {
		r := true
		for m, _ := range g.edges {
			if n == m {
				r = false
			}
		}
		if r {
			return n
		}
	}

	return nil
}

// Setup ...
func (g *Graph) Setup(in *os.File) error {
	g.nodes = map[string]*Node{}
	g.edges = map[*Node][]*Node{}

	err := lib.LoopOverLines(in, func(line []byte) error {
		bits := strings.Split(string(line), " ")
		end := bits[1]
		start := bits[7]

		a := g.addNode(start)
		b := g.addNode(end)
		a.AddRequirement(b)
		g.addEdge(b, a)

		return nil
	})

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Children ...
func (g Graph) Children(n *Node) []*Node {
	out, ok := g.edges[n]
	if !ok {
		return []*Node{}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

// FirstNode ...
func (g Graph) FirstNodes() []*Node {
	var first []*Node

	areChildren := []*Node{}

	for _, node := range g.nodes {
		connections := g.edges[node]

		for _, child := range connections {
			add := true
			for _, x := range areChildren {
				if x == child {
					add = false
				}
			}
			if add {
				areChildren = append(areChildren, child)
			}
		}
	}

	for _, node := range g.nodes {
		// spew.Dump(node)
		isChild := false
		for _, n := range areChildren {
			if n == node {
				isChild = true
			}
		}
		if !isChild {
			first = append(first, node)
		}
	}
	sort.Slice(first, func(i, j int) bool { return first[i].Name < first[j].Name })
	return first
}

// Print ...
func (g Graph) Print() string {
	f := g.FirstNodes()

	fmt.Printf("have %v nodes\n", len(g.nodes))

	elements := []*Node{}
	elements = append(elements, f...)

	for i := 0; i < len(elements); i++ {
		x := i + 1
		n := elements[i]
		children := g.Children(n)

		for _, m := range children {
			met := m.MeetsRequirements(elements[:x])
			added := false

			for _, z := range elements {
				if m == z {
					added = true
				}
			}

			if met && !added {
				elements = append(elements, &Node{})
				copy(elements[x+1:], elements[x:])
				elements[x] = m
				x++
			}
		}
	}

	out := ""
	for _, n := range elements {
		out = fmt.Sprintf("%v%v", out, n.Name)
	}

	fmt.Printf("chain has %v elements\n", len(out))
	return out
}

// AddNode ...
func (g *Graph) addNode(a string, requires ...string) *Node {
	n := g.FindNode(a)
	if n != nil {
		return n
	}

	n = &Node{Name: a, Requires: []*Node{}}
	g.nodes[a] = n
	return n
}

// AddEdge ...
func (g *Graph) addEdge(from, to *Node) {
	bits, ok := g.edges[from]
	if !ok {
		bits = []*Node{}
	}
	bits = append(bits, to)
	g.edges[from] = bits
}

// FindNode ...
func (g Graph) FindNode(f string) *Node {
	n, ok := g.nodes[f]
	if ok {
		return n
	}
	return nil
}
