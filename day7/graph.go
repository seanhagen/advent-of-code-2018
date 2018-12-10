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

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Graph data structure
type Graph struct {
	nodes map[string]*Node
	edges map[*Node][]*Node

	done      []*Node
	needsWork []*Node
	locked    []*Node

	workers map[int]*worker

	baseWorkTime int
	numWorkers   int
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
	out = sortnodes(out)
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
	first = sortnodes(first)
	return first
}

func (g Graph) Print() string {
	f := g.FirstNodes()

	elements := []*Node{}

	toAdd := []*Node{}
	toAdd = append(toAdd, f...)

	toAdd = sortnodes(toAdd)

	for i := 0; i < len(toAdd); i++ {
		n := toAdd[i]
		met := n.MeetsRequirements(elements)
		if met {
			// add node to elements
			found := false
			for _, m := range elements {
				if m == n {
					found = true
				}
			}
			if !found {
				elements = append(elements, n)
			}

			// remove node from toAdd
			copy(toAdd[i:], toAdd[i+1:])
			toAdd[len(toAdd)-1] = nil // or the zero value of T
			toAdd = toAdd[:len(toAdd)-1]

			// fetch nodes children
			children := g.Children(n)
			// add children to toAdd
			toAdd = append(toAdd, children...)

			// sort toAdd by node name alphabetically
			toAdd = sortnodes(toAdd)

			// start over at the beginning
			i = -1
		}

		if len(toAdd) == 0 {
			break
		}
	}

	out := ""
	for _, n := range elements {
		out = fmt.Sprintf("%v%v", out, n.Name)
	}

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

func sortnodes(in []*Node) []*Node {
	sort.Slice(in, func(i, j int) bool { return in[i].Name < in[j].Name })
	return in
}
