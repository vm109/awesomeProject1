package graphs

import (
	"fmt"
	"testing"
)

func TestPrint_Graphs(t *testing.T) {
	Print_Graphs()
}

/*
we have created a two directional Adjancey Graph using AjacentMatrixFrom Edges Array
 */
func TestCreateAdjacenyMatricFromEdges(t *testing.T) {
	edges_arr := []GraphEdge{
		GraphEdge{StartNode: 0, EndNode: 4},
		GraphEdge{StartNode: 1, EndNode: 4},
		GraphEdge{StartNode: 0, EndNode: 1},
		GraphEdge{StartNode: 3, EndNode: 4},
		GraphEdge{StartNode: 1, EndNode: 3},
		GraphEdge{StartNode: 1, EndNode: 2},
		GraphEdge{StartNode: 3, EndNode: 2},
	}
	graph_adjacency_map := CreateAdjacenyMatricFromEdges(edges_arr)
	for key, value := range graph_adjacency_map{
		fmt.Println("printing adjacent nodes of ", key)
		for _, graphNodeValue := range value{
			dest_value := graphNodeValue.Data
			fmt.Println("there is a edge between ", key, " and ", dest_value)
		}
	}
}
