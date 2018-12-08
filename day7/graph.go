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
}

// Start ...
func (n *Node) Start(s, base int) {
	// fmt.Printf("node starting: %v\n", n.Name)
	for i := range letters {
		l := string(letters[i])
		if l == n.Name {
			n.time = i + 1
		}
	}

	n.start = s
	n.end = s + n.time
	n.working = true

	fmt.Printf("node %v start %v end %v\n", n.Name, n.start, n.end)
}

// Done ...
func (n *Node) Done(s int) bool {
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

// Graph data structure
type Graph struct {
	nodes map[string]*Node
	edges map[*Node][]*Node

	unstarted  []*Node
	inProgress []*Node
	done       []*Node
	work       map[int][]int

	baseWorkTime int
	numWorkers   int
	finishTime   int
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

func (g Graph) Print() string {
	f := g.FirstNodes()

	elements := []*Node{}

	toAdd := []*Node{}
	toAdd = append(toAdd, f...)

	sort.Slice(toAdd, func(i, j int) bool { return toAdd[i].Name < toAdd[j].Name })

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
			sort.Slice(toAdd, func(i, j int) bool { return toAdd[i].Name < toAdd[j].Name })

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

func (g *Graph) SetupWork(workers, base int) {
	g.numWorkers = workers
	g.baseWorkTime = base

	start := g.FirstNodes()

	g.unstarted = append(g.unstarted, start...)
}

// DoWork ...
func (g *Graph) DoWork() {
	i := 0
	for ; g.workLeft(); i++ {
		fmt.Printf("second: %v\n", i)

		for j := 0; j < len(g.inProgress); j++ {
			n := g.inProgress[j]
			if n.Done(i) {
				fmt.Printf("node %v is done work! (sec: %v)\n", n.Name, i)
				// spew.Dump(n)
				children := g.Children(n)
				fmt.Printf("node %v has %v children: ", n.Name, len(children))
				for _, x := range children {
					fmt.Printf("%v ", x.Name)
				}
				fmt.Printf("\n")

				for _, x := range children {
					found := false

					for _, y := range g.unstarted {
						if y == x {
							found = true
						}
					}

					if !found {
						g.unstarted = append(g.unstarted, x)
					}
				}

				//g.unstarted = append(g.unstarted, children...)

				fmt.Printf("unstarted now contains: ")
				for _, x := range g.unstarted {
					fmt.Printf("%v ", x.Name)
				}
				fmt.Printf("\n")

				sort.Slice(g.unstarted, func(i, j int) bool { return g.unstarted[i].Name < g.unstarted[j].Name })

				g.done = append(g.done, n)

				fmt.Printf("gotta remove node %v from in progress\n", n.Name)
				fmt.Printf("currently in progress (len: %v): ", len(g.inProgress))
				for _, x := range g.inProgress {
					fmt.Printf("%v ", x.Name)
				}
				fmt.Printf("\n")

				if len(g.inProgress) == 1 {
					g.inProgress = []*Node{}
				} else {
					if i == len(g.inProgress)-1 {
						g.inProgress = g.inProgress[:i-1]
					} else {
						fmt.Printf("copy: len: %v, i: %v\n", len(g.inProgress), j)
						copy(g.inProgress[j:], g.inProgress[j+1:])
						g.inProgress[len(g.inProgress)-1] = nil // or the zero value of T
						g.inProgress = g.inProgress[:len(g.inProgress)-1]
					}
				}
			}

		}

		if len(g.unstarted) > 0 {
			toAdd := g.numWorkers - len(g.inProgress)
			fmt.Printf("need to add %v jobs (len unstarted: %v)\n", toAdd, len(g.unstarted))
			for j := 0; j < toAdd && j < len(g.unstarted); j++ {
				n := g.unstarted[j]
				n.Start(i, g.baseWorkTime)
				fmt.Printf("adding node to be worked on: %v (start: %v, will end at: %v)\n", n.Name, n.start, n.end)
				g.inProgress = append(g.inProgress, n)

				copy(g.unstarted[j:], g.unstarted[j+1:])
				g.unstarted[len(g.unstarted)-1] = nil // or the zero value of T
				g.unstarted = g.unstarted[:len(g.unstarted)-1]
				j = -1
			}
		}

		// spew.Dump(g.unstarted, g.inProgress)
		// break

		fmt.Printf("------\n\n")

		if len(g.done) == 26 {
			break
		}

		if i > 1000 {
			break
		}

	}

	g.finishTime = i
	fmt.Printf("work done\n")
}

// WorkDone ...
func (g *Graph) workLeft() bool {
	return len(g.unstarted) > 0 || len(g.inProgress) > 0
}

// PrintWork ...
func (g *Graph) PrintWork() string {
	return ""
}

// WorkTime ...
func (g *Graph) WorkTime() int {
	return g.finishTime
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
