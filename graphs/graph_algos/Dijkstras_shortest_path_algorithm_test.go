package graph_algos

import (
	"awesomeProject1/my_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Single Source Shortest Path - Shortest path in a graph is found using a
single vertex as the source vertex
[1,2,3,4,5,6]
[0,1,2,3,4,5]
1 -> 2 [weight 2]
1 -> 3 [weight 4]
2 -> 3 [weight 1]
2 -> 4 [weight 7]
3 -> 5 [weight 3]
4 -> 6 [weight 1]
5 -> 4 [weight 2]
5 -> 6 [weight 5]
 */
func TestDijkstras_Find_Shortest_Path(t *testing.T) {
//	nodes_list := []int32{ 1, 2 ,3, 4,5 ,6}
	edge_list := []my_utils.IndexGraphEdge{}
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 1, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 2, Weight: 4})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 2, Weight: 1})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 3, Weight: 7})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 2, End: 4, Weight: 3})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 3, End: 5, Weight: 1})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 3, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 5, Weight: 5})
	source_vertex := 4
	distance, nodes_visited := Dijkstras_Find_Shortest_Path(edge_list, 6, int32(source_vertex)  )
	assert.Equal(t, distance, int32(3))
	assert.Equal(t, nodes_visited[0], int32(4))
	assert.Equal(t, nodes_visited[1], int32(3))
	assert.Equal(t, nodes_visited[2], int32(5))
}

func TestDijkstras_Find_Shortest_Path_Starting_at_0(t *testing.T) {
	//	nodes_list := []int32{ 1, 2 ,3, 4,5 ,6}
	edge_list := []my_utils.IndexGraphEdge{}
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 1, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 2, Weight: 4})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 2, Weight: 1})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 3, Weight: 7})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 2, End: 4, Weight: 3})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 3, End: 5, Weight: 1})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 3, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 5, Weight: 5})
	source_vertex := 0
	distance, nodes_visited := Dijkstras_Find_Shortest_Path(edge_list, 6, int32(source_vertex)  )
	assert.Equal(t, distance, int32(9))
	assert.Equal(t, nodes_visited[0], int32(0))
	assert.Equal(t, nodes_visited[1], int32(1))
	assert.Equal(t, nodes_visited[2], int32(2))
	assert.Equal(t, nodes_visited[3], int32(4))
	assert.Equal(t, nodes_visited[4], int32(3))
	assert.Equal(t, nodes_visited[5], int32(5))

}
