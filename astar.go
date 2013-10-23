// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package astar implements the A* shortest path finding algorithm.
package astar

import (
	"container/heap"
)

type Node interface{}

type Graph interface {
	Neighbours(n Node) []Node
}

type CostFunc func(a, b Node) float64

type Path []Node

func newPath(start Node) Path {
	return []Node{start}
}

func (p Path) last() Node {
	return p[len(p)-1]
}

func (p Path) cont(n Node) Path {
	newPath := make([]Node, len(p), len(p)+1)
	copy(newPath, p)
	newPath = append(newPath, n)
	return newPath
}

func (p Path) Cost(d CostFunc) (c float64) {
	for i := 1; i < len(p); i++ {
		c += d(p[i-1], p[i])
	}
	return c
}

// FindPath finds the shortest path between start and dest in graph g
// using the cost function d and the cost heuristic function h.
func FindPath(g Graph, start, dest Node, d, h CostFunc) Path {
	closed := make(map[Node]bool)

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{value: newPath(start)})

	for pq.Len() > 0 {
		p := heap.Pop(pq).(*item).value.(Path)
		if closed[p.last()] {
			continue
		}
		if p.last() == dest {
			// Path found
			return p
		}
		n := p.last()
		closed[n] = true

		for _, nb := range g.Neighbours(n) {
			newPath := p.cont(nb)
			heap.Push(pq, &item{
				value:    newPath,
				priority: newPath.Cost(d) + h(nb, dest),
			})
		}
	}

	// No path found
	return nil
}
