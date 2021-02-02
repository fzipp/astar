package astar_test

import (
	"fmt"
	"image"
	"math"

	"github.com/fzipp/astar"
)

func ExampleFindPath_maze() {
	maze := floorPlan{
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
	path := astar.FindPath(maze, start, dest, distance, distance)

	// Mark the path with dots before printing
	for _, p := range path {
		maze.put(p.(image.Point), '.')
	}
	maze.print()
	// Output:
	// ###############
	// #   # #     #.#
	// # ### ### ###.#
	// #   # # #   #.#
	// ### # # # ###.#
	// # # #  .......#
	// # # ###.### ###
	// #   # #.# #   #
	// ### # #.# # ###
	// # #.....  # # #
	// # #.######### #
	// #...      #   #
	// #.### # # ### #
	// #.  # # #     #
	// ###############
}

// distance is our cost function. We use points as nodes, so we
// calculate their Euclidean distance.
func distance(a, b astar.Node) float64 {
	p := a.(image.Point)
	q := b.(image.Point)
	d := q.Sub(p)
	return math.Sqrt(float64(d.X*d.X + d.Y*d.Y))
}

type floorPlan []string

// Neighbours implements the astar.Graph interface
func (f floorPlan) Neighbours(n astar.Node) []astar.Node {
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
		if f.isFreeAt(q) {
			res = append(res, q)
		}
	}
	return res
}

func (f floorPlan) isFreeAt(p image.Point) bool {
	return f.isInBounds(p) && f[p.Y][p.X] == ' '
}

func (f floorPlan) isInBounds(p image.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(f) && p.X < len(f[p.Y])
}

func (f floorPlan) put(p image.Point, c rune) {
	f[p.Y] = f[p.Y][:p.X] + string(c) + f[p.Y][p.X+1:]
}

func (f floorPlan) print() {
	for _, row := range f {
		fmt.Println(row)
	}
}
