package graph_cycle

import (
	"fmt"
	"github.com/myawesomeproject1/graphs/graph_interfaces"
	"testing"
)

func TestAdjacentMatrix_FindGraphCycle(t *testing.T) {
	var graph graph_interfaces.Graph = NewAdjacentList(3)
	graph.AddEdgeIntoAdjacentList(0, 1)
	graph.AddEdgeIntoAdjacentList(0, 2)
	graph.AddEdgeIntoAdjacentList(1, 2)
	fmt.Println(graph)
	hasCycle := graph.FindGraphCycle()
	fmt.Println(hasCycle)
}
