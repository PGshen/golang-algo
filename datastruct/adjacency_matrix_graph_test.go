package datastruct

import (
	"fmt"
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
	graph.addEdge(1, 5)
	graph.addEdge(0, 1)
	graph.print()

	// graph.removeVertex(2)
	graph.print()

	fmt.Printf("%v\n", graph.bfs())
	fmt.Printf("%v\n", graph.dfs())
}
