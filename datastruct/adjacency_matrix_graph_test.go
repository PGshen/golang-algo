package datastruct

import (
	"testing"
)

func TestAdjMatGraph(t *testing.T) {
	vertices := []int{1, 2, 3, 4, 5}
	edges := [][]int{
		{1, 2},
		{2, 4},
	}
	graph := newAdjMatGraph(vertices, edges)
	graph.print()

	graph.addEdge(0, 3)
	graph.addEdge(2, 3)
	graph.print()

	graph.removeEdge(1, 2)
	graph.print()

	graph.addVertex(10)
	graph.print()

	graph.removeVertex(2)
	graph.print()
}
