package day7

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/seanhagen/advent-of-code-2018/lib"
)

// https://flaviocopes.com/golang-data-structure-graph/
// https://godoc.org/github.com/gonum/graph

// Node point in graph
type Node struct {
	name string
}

// Graph data structure
type Graph struct {
	nodes map[string]*Node
}

// Setup ...
func (g *Graph) Setup(in *os.File) error {
	g.nodes = map[string]*Node{}

	err := lib.LoopOverLines(in, func(line []byte) error {
		// fmt.Printf("%v\n", string(line))
		bits := strings.Split(string(line), " ")
		// spew.Dump(bits)

		end := bits[1]
		start := bits[7]

		fmt.Printf("%v -> %v\n", start, end)

		return nil
	})

	if err != nil && err != io.EOF {
		return err
	}

	return err
}

// Print ...
func (g Graph) Print() string {
	out := ""
	return out
}

// AddNode ...
func (g *Graph) addNode(a string) *Node {
	n := g.findNode(a)
	if n != nil {
		return n
	}

	n = &Node{a}
	g.nodes[a] = n
	return n
}

// AddEdge ...
func (g *Graph) addEdge(n, m string) {

}

// findNode ...
func (g Graph) findNode(f string) *Node {
	n, ok := g.nodes[f]
	if ok {
		return n
	}
	return nil
}
