package traversals

import (
	"awesomeProject1/graphs"
	"fmt"
	"testing"
)

func TestDFS_Graph(t *testing.T) {
	edges := []*graphs.GraphEdge{
		&graphs.GraphEdge{0, 1},
		&graphs.GraphEdge{0,4},
		&graphs.GraphEdge{1,2},
		&graphs.GraphEdge{1,3},
		&graphs.GraphEdge{1,4},
		&graphs.GraphEdge{2,3},
		&graphs.GraphEdge{3, 4},
	}
	graph_map := graphs.Represent_Graph_Using_Map(edges)

	dfs_map := DFS_Graph(graph_map)
	for key, _ := range dfs_map {
		fmt.Println(key)
	}
}
