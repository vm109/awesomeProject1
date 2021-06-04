package graphs

import (
	"fmt"
	"testing"
)

func TestDFS_Graph(t *testing.T) {
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

	dfs_map := DFS_Graph(graph_map)
	for key, _ := range dfs_map {
		fmt.Println(key)
	}
}
