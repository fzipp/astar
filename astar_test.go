// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar

import (
	"image"
	"math"
	"testing"
)

type graph map[Node][]Node

func newGraph() graph {
	return make(map[Node][]Node)
}

func (g graph) Link(a, b Node) graph {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}

func (g graph) Neighbours(n Node) []Node {
	return g[n]
}

func nodeDist(a, b Node) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

func TestFindPath(t *testing.T) {
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)
	g := newGraph().Link(a, b).Link(a, c).Link(b, d).Link(c, d)

	want := Path{
		image.Pt(2, 3),
		image.Pt(1, 6),
		image.Pt(5, 6),
	}

	p := FindPath(g, a, d, nodeDist, nodeDist)
	if len(p) != len(want) {
		t.Errorf("Returned path has %d nodes, want %d nodes.", len(p), len(want))
	}
	for i, n := range p {
		if n != want[i] {
			t.Errorf("Node %d of path is %v, want %v.", i, n, want[i])
		}
	}
}
