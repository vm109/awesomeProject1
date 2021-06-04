package graphs

import (
	"fmt"
	"testing"
)

func TestRepresent_Graph_Using_Map(t *testing.T) {
	edges := []*GraphEdge{
		&GraphEdge{ 0, 1},
		&GraphEdge{0,4},
		&GraphEdge{1,2},
		&GraphEdge{1,3},
		&GraphEdge{1,4},
		&GraphEdge{2,3},
		&GraphEdge{3, 4},
	}
	graph_map :=Represent_Graph_Using_Map(edges)
	for key, val := range graph_map {
		fmt.Println(key)
		fmt.Println(val)
	}
}