package astar_test

import (
	"fmt"
	"image"
	"math"

	"github.com/fzipp/astar"
)

type graph map[astar.Node][]astar.Node

func newGraph() graph {
	return make(map[astar.Node][]astar.Node)
}

func (g graph) Link(a, b astar.Node) graph {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}

func (g graph) Neighbours(n astar.Node) []astar.Node {
	return g[n]
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their euclidian distance.
func nodeDist(a, b astar.Node) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

func ExampleFindPath() {
	// Create a graph with points as nodes
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)
	g := newGraph().Link(a, b).Link(a, c).Link(b, c).Link(b, d).Link(c, d)

	// Find the shortest path from a to d
	p := astar.FindPath(g, a, d, nodeDist, nodeDist)

	// Output the result
	if p == nil {
		fmt.Println("No path found.")
		return
	}
	for i, n := range p {
		fmt.Printf("%d: %s\n", i, n)
	}
	// Output:
	// 0: (2,3)
	// 1: (1,6)
	// 2: (5,6)
}
