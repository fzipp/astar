// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar

import (
	"container/heap"
	"testing"
)

func TestPushPop(t *testing.T) {
	pq := &priorityQueue{}
	heap.Init(pq)

	want := "ebdac"
	heap.Push(pq, &item{value: "a", priority: 1.2})
	heap.Push(pq, &item{value: "b", priority: 5})
	heap.Push(pq, &item{value: "c", priority: -0.4})
	heap.Push(pq, &item{value: "d", priority: 3.7})
	heap.Push(pq, &item{value: "e", priority: 11})

	s := ""
	for pq.Len() > 0 {
		s += heap.Pop(pq).(*item).value.(string)
	}

	if s != want {
		t.Errorf("Retrieved item order was %q, want %q.", s, want)
	}
}
