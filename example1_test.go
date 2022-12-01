// Copyright 2018 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar_test

import (
	"fmt"
	"image"
	"math"

	"github.com/fzipp/astar"
)

func ExampleFindPath() {
	// Create a graph with 2D points as nodes
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)
	g := newGraph[image.Point]().link(a, b).link(a, c).link(b, c).link(b, d).link(c, d)

	// Find the shortest path from a to d
	p := astar.FindPath[image.Point](g, a, d, nodeDist, nodeDist)

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

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

// graph is represented by an adjacency list.
type graph[Node comparable] map[Node][]Node

func newGraph[Node comparable]() graph[Node] {
	return make(map[Node][]Node)
}

// link creates a bi-directed edge between nodes a and b.
func (g graph[Node]) link(a, b Node) graph[Node] {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}

// Neighbours returns the neighbour nodes of node n in the graph.
func (g graph[Node]) Neighbours(n Node) []Node {
	return g[n]
}
