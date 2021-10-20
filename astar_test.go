// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar

import (
	"image"
	"math"
	"testing"
)

type graph[Node comparable] map[Node][]Node

func newGraph[Node comparable]() graph[Node] {
	return make(map[Node][]Node)
}

func (g graph[Node]) Link(a, b Node) graph[Node] {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}

func (g graph[Node]) Neighbours(n Node) []Node {
	return g[n]
}

func nodeDist(p, q image.Point) float64 {
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

func TestFindPath(t *testing.T) {
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)
	g := newGraph[image.Point]().Link(a, b).Link(a, c).Link(b, d).Link(c, d)

	want := Path[image.Point]{
		image.Pt(2, 3),
		image.Pt(1, 6),
		image.Pt(5, 6),
	}

	p := FindPath[image.Point](g, a, d, nodeDist, nodeDist)
	if len(p) != len(want) {
		t.Errorf("Returned path has %d nodes, want %d nodes.", len(p), len(want))
	}
	for i, n := range p {
		if n != want[i] {
			t.Errorf("Node %d of path is %v, want %v.", i, n, want[i])
		}
	}
}
