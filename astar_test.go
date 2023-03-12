// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar_test

import (
	"image"
	"reflect"
	"testing"

	"github.com/fzipp/astar"
)

func TestFindPath(t *testing.T) {
	a := image.Pt(2, 3)
	b := image.Pt(1, 7)
	c := image.Pt(1, 6)
	d := image.Pt(5, 6)

	tests := []struct {
		name  string
		graph graph[image.Point]
		start image.Point
		dest  image.Point
		want  astar.Path[image.Point]
	}{
		{
			name: "find simple path",
			graph: newGraph[image.Point]().
				link(a, b).link(a, c).
				link(b, d).
				link(c, d),
			start: a,
			dest:  d,
			want: astar.Path[image.Point]{
				image.Pt(2, 3),
				image.Pt(1, 6),
				image.Pt(5, 6),
			},
		},
		{
			name: "find no path",
			graph: newGraph[image.Point]().
				link(a, b).link(a, c).
				link(b, d).
				link(c, d),
			start: a,
			dest:  image.Pt(1, 1),
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := astar.FindPath[image.Point](tt.graph, tt.start, tt.dest, nodeDist, nodeDist)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngraph = %v\nFindPath(graph, %v, %v, nodeDist, nodeDist) = %v, want %v",
					tt.graph, tt.start, tt.dest, got, tt.want)
			}
		})
	}
}
