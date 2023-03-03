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
	p1 := image.Pt(3, 1)
	p2 := image.Pt(1, 2)
	p3 := image.Pt(2, 4)
	p4 := image.Pt(4, 5)
	p5 := image.Pt(4, 3)
	p6 := image.Pt(5, 1)
	p7 := image.Pt(8, 4)
	p8 := image.Pt(8, 3)
	p9 := image.Pt(6, 3)
	g := newGraph[image.Point]().
		link(p1, p2).link(p1, p3).
		link(p2, p3).link(p2, p4).
		link(p3, p4).link(p3, p5).
		link(p4, p6).link(p4, p7).
		link(p5, p7).
		link(p6, p9).
		link(p7, p8).
		link(p8, p9)

	// Find the shortest path from p1 to p9
	p := astar.FindPath[image.Point](g, p1, p9, nodeDist, nodeDist)

	// Output the result
	if p == nil {
		fmt.Println("No path found.")
		return
	}
	for i, n := range p {
		fmt.Printf("%d: %s\n", i, n)
	}
	// Output:
	// 0: (3,1)
	// 1: (2,4)
	// 2: (4,5)
	// 3: (5,1)
	// 4: (6,3)
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
