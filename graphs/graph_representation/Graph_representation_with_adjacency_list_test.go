package graph_representation

import (
	"fmt"
	"testing"
)

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
