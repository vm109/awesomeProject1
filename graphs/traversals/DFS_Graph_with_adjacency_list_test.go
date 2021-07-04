package traversals

import (
	"awesomeProject1/graphs"
	"awesomeProject1/graphs/graph_representation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildTestGraphAdjacenyList() GraphRepresentedAsAdjacencyList{
	graphEdgesList := make([]graphs.GraphEdge, 0)
	edge1:= graphs.GraphEdge{StartNode: 0, EndNode: 1}
	edge2:= graphs.GraphEdge{StartNode: 0, EndNode: 2}
	edge3:= graphs.GraphEdge{StartNode: 1, EndNode: 2}
	edge4:= graphs.GraphEdge{StartNode: 2, EndNode: 3}
	edge5:= graphs.GraphEdge{StartNode: 2, EndNode: 0}
	edge6:= graphs.GraphEdge{StartNode: 3, EndNode: 3}
	graphEdgesList = append(graphEdgesList, edge1)
	graphEdgesList = append(graphEdgesList, edge2)
	graphEdgesList = append(graphEdgesList, edge3)
	graphEdgesList = append(graphEdgesList, edge4)
	graphEdgesList = append(graphEdgesList, edge5)
	graphEdgesList = append(graphEdgesList, edge6)
	return GraphRepresentedAsAdjacencyList{
		graph_representation.RepresentGraphAsAdjacencyListWhenEdgesListIsGiven(graphEdgesList),
	}
}
func TestGraphRepresentedAsAdjacencyList_DFS_Graph_with_adjacency_list(t *testing.T) {
	visited_nodes := make(map[int]bool, 0)
	final_dfs_visited_nodes := buildTestGraphAdjacenyList().DFS_Graph_with_adjacency_list(2, visited_nodes)
	assert.Equal(t, final_dfs_visited_nodes, []int{2, 3, 0, 1})
}


func TestGraphRepresentedAsAdjacencyList_DFS_Graph_with_adjacency_list_Start_with_3(t *testing.T) {
	visited_nodes := make(map[int]bool, 0)
	final_dfs_visited_nodes := buildTestGraphAdjacenyList().DFS_Graph_with_adjacency_list(3, visited_nodes)
	assert.Equal(t, final_dfs_visited_nodes, []int{3})
}

func TestGraphRepresentedAsAdjacencyList_DFS_Graph_with_adjacency_list_start_with_1(t *testing.T) {
	visited_nodes := make(map[int]bool, 0)
	final_dfs_visited_nodes := buildTestGraphAdjacenyList().DFS_Graph_with_adjacency_list(1, visited_nodes)
	assert.Equal(t, final_dfs_visited_nodes, []int{1, 2, 3, 0})
}

func TestGraphRepresentedAsAdjacencyList_DFS_Graph_with_adjacency_list_start_with_0(t *testing.T) {
	visited_nodes := make(map[int]bool, 0)
	final_dfs_visited_nodes := buildTestGraphAdjacenyList().DFS_Graph_with_adjacency_list(0, visited_nodes)
	assert.Equal(t, final_dfs_visited_nodes, []int{0, 1, 2, 3})
}