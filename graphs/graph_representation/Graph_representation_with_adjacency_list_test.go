package graph_representation

import (
	"fmt"
	"github.com/myawesomeproject1/graphs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepresentGraphAsAdjacencyListWhenEdgesListIsGiven(t *testing.T) {
	graphEdgesList := make([]graphs.GraphEdge, 0)
	edge1 := graphs.GraphEdge{StartNode: 0, EndNode: 1}
	edge2 := graphs.GraphEdge{StartNode: 0, EndNode: 2}
	edge3 := graphs.GraphEdge{StartNode: 1, EndNode: 2}
	edge4 := graphs.GraphEdge{StartNode: 2, EndNode: 3}
	edge5 := graphs.GraphEdge{StartNode: 2, EndNode: 0}
	edge6 := graphs.GraphEdge{StartNode: 3, EndNode: 3}
	graphEdgesList = append(graphEdgesList, edge1)
	graphEdgesList = append(graphEdgesList, edge2)
	graphEdgesList = append(graphEdgesList, edge3)
	graphEdgesList = append(graphEdgesList, edge4)
	graphEdgesList = append(graphEdgesList, edge5)
	graphEdgesList = append(graphEdgesList, edge6)
	graph_adjacency_list := RepresentGraphAsAdjacencyListWhenEdgesListIsGiven(graphEdgesList)
	for key, adjacent_list := range graph_adjacency_list {
		if key == 2 {
			assert.Equal(t, len(adjacent_list), 2)
		}
		if key == 3 {
			assert.Equal(t, len(adjacent_list), 1)
		}
		if key == 0 {
			assert.Equal(t, len(adjacent_list), 2)
		}
		if key == 1 {
			assert.Equal(t, len(adjacent_list), 1)
		}
	}
}

/*
Adjacency List Representation with array of array. Each array holds the adjacent nodes to the Array index value.

0 -> 1
0 -> 4
1 -> 4
1 -> 3
1 -> 2
2 -> 3
3 -> 4
*/
func TestAdjacencyListRepAddition(t *testing.T) {
	adjList := NewAdjacencyListRep(5)
	adjList.AddEdgeIntoAdjacentList(0, 1)
	adjList.AddEdgeIntoAdjacentList(0, 4)
	adjList.AddEdgeIntoAdjacentList(1, 4)
	adjList.AddEdgeIntoAdjacentList(1, 3)
	adjList.AddEdgeIntoAdjacentList(1, 2)
	adjList.AddEdgeIntoAdjacentList(2, 3)
	adjList.AddEdgeIntoAdjacentList(3, 4)
	for _, list := range adjList {
		for _, val := range list {
			fmt.Print(val, ",")
		}
		fmt.Println()
	}
}
