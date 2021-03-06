# astar

[![PkgGoDev](https://pkg.go.dev/badge/github.com/fzipp/astar)](https://pkg.go.dev/github.com/fzipp/astar)
![Build Status](https://github.com/fzipp/astar/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/fzipp/astar)](https://goreportcard.com/report/github.com/fzipp/astar)

Package astar implements the
[A* shortest path finding algorithm](https://en.wikipedia.org/wiki/A*_search_algorithm).

## Examples

In order to use the `astar.FindPath` function to find the shortest path
between two nodes of a graph you need a graph data structure that implements
the `Neighbours` method to  satisfy the `astar.Graph` interface, and a cost
function. It is up to you how the graph is internally implemented.

### A maze

In this example the graph is represented by a slice of strings, each character
representing a cell of a floor plan. Graph nodes are cell positions
as `image.Point` values, with (0, 0) at the upper left corner. 
Spaces represent free cells available for walking, other characters like
`#` represent walls.
The `Neighbours` method returns the positions of the adjacent free cells
to the north, east, south, and west of a given position (diagonal movement
is not allowed in this example).

The cost function `nodeDist` simply calculates the Euclidean distance
between two cell positions.

```go
package main

import (
	"fmt"
	"image"
	"math"

	"github.com/fzipp/astar"
)

func main() {
	maze := graph{
		"###############",
		"#   # #     # #",
		"# ### ### ### #",
		"#   # # #   # #",
		"### # # # ### #",
		"# # #         #",
		"# # ### ### ###",
		"#   # # # #   #",
		"### # # # # ###",
		"# #       # # #",
		"# # ######### #",
		"#         #   #",
		"# ### # # ### #",
		"#   # # #     #",
		"###############",
	}
	start := image.Pt(1, 13) // Bottom left corner
	dest := image.Pt(13, 1)  // Top right corner

	// Find the shortest path
	path := astar.FindPath(maze, start, dest, nodeDist, nodeDist)
	
	// Mark the path with dots before printing
	for _, p := range path {
		maze.put(p.(image.Point), '.')
	}
	maze.print()
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func nodeDist(a, b astar.Node) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type graph []string

// Neighbours implements the astar.Graph interface
func (g graph) Neighbours(n astar.Node) []astar.Node {
	p := n.(image.Point)
	offsets := []image.Point{
		image.Pt(0, -1), // North
		image.Pt(1, 0),  // East
		image.Pt(0, 1),  // South
		image.Pt(-1, 0), // West
	}
	res := make([]astar.Node, 0, 4)
	for _, off := range offsets {
		q := p.Add(off)
		if g.isFreeAt(q) {
			res = append(res, q)
		}
	}
	return res
}

func (g graph) isFreeAt(p image.Point) bool {
	return g.isInBounds(p) && g[p.Y][p.X] == ' '
}

func (g graph) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}

func (g graph) put(p image.Point, c rune) {
	g[p.Y] = g[p.Y][:p.X] + string(c) + g[p.Y][p.X+1:]
}

func (g graph) print() {
	for _, row := range g {
		fmt.Println(row)
	}
}
```

Output:

```
###############
#   # #     #.#
# ### ### ###.#
#   # # #   #.#
### # # # ###.#
# # #  .......#
# # ###.### ###
#   # #.# #   #
### # #.# # ###
# #.....  # # #
# #.######### #
#...      #   #
#.### # # ### #
#.  # # #     #
###############
```

### 2D points as nodes

In this example the graph is represented by an adjacency list. Nodes are
2D points in Euclidean space as `image.Point` values. The `link` function
creates a bi-directed edge between a pair of nodes.

The cost function `nodeDist` calculates the Euclidean distance
between two points (nodes).

```go
package main

import (
	"fmt"
	"image"
	"math"

	"github.com/fzipp/astar"
)

func main() { 
	// Create a graph with 2D points as nodes
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)
	g := newGraph().link(a, b).link(a, c).link(b, c).link(b, d).link(c, d)

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
}

// graph is represented by an adjacency list.
type graph map[astar.Node][]astar.Node

func newGraph() graph {
	return make(map[astar.Node][]astar.Node)
}

// link creates a bi-directed edge between nodes a and b.
func (g graph) link(a, b astar.Node) graph {
	g[a] = append(g[a], b)
	g[b] = append(g[b], a)
	return g
}

// Neighbours returns the neighbour nodes of node n in the graph.
func (g graph) Neighbours(n astar.Node) []astar.Node {
	return g[n]
}

// nodeDist is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func nodeDist(a, b astar.Node) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}
```

Output:

```
0: (2,3)
1: (1,6)
2: (5,6)
```

## License

This project is free and open source software licensed under the
[BSD 3-Clause License](LICENSE).
