package traversals

import (
	"awesomeProject1/graphs"
	"testing"
)

func TestBFS_graph(t *testing.T) {
	edges := []*graphs.GraphEdge{
		&graphs.GraphEdge{1,2},
		&graphs.GraphEdge{1, 3},
		&graphs.GraphEdge{2, 4},
		&graphs.GraphEdge{4, 2},
		&graphs.GraphEdge{3, 5},
		&graphs.GraphEdge{5, 1},
		&graphs.GraphEdge{5, 1},
	}
	constructed_graph := graphs.Represent_Graph_Using_Map(edges)
	BFS_graph(&constructed_graph)
}
