package graph_representation

import (
	"awesomeProject1/my_utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepresentGraphAsMatrixWhenEdgeListIsGiven(t *testing.T) {
	edge_list := []*my_utils.IndexGraphEdge{}
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 1, Weight: 5})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 2, Weight: 7})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 3, Weight: 3})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 1, End: 4, Weight: 2})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 1, End: 5, Weight: 10})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 2, End: 6, Weight: 1})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 3, End: 7, Weight: 11})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 4, End: 7, Weight: 9})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 5, End: 7, Weight: 4})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 6, End: 7, Weight: 6})
	vertices := RepresentGraphAsMatrixWhenEdgeListIsGiven(edge_list, 8)
	assert.Equal(t, []int{0,0,0,1,1,1,1,0}, vertices[7])
}

func TestFindTheAdjacentNodesOfGvenNode(t *testing.T){
	edge_list := []*my_utils.IndexGraphEdge{}
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 1, Weight: 5})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 2, Weight: 7})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 0, End: 3, Weight: 3})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 1, End: 4, Weight: 2})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 1, End: 5, Weight: 10})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 2, End: 6, Weight: 1})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 3, End: 7, Weight: 11})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 4, End: 7, Weight: 9})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 5, End: 7, Weight: 4})
	edge_list = append(edge_list, &my_utils.IndexGraphEdge{Start: 6, End: 7, Weight: 6})
	vertices := RepresentGraphAsMatrixWhenEdgeListIsGiven(edge_list, 8)
	fmt.Println(vertices)
	// find the adjacent nodes of 2 - expecting 0 and 6 as the adjacent vertices
	assert.Equal(t, 1,vertices[2][0])
	assert.Equal(t, 1, vertices[2][6])
}


func TestRepresentDirectedGraphAsWeightedAdjacentMatrix(t *testing.T){
	edge_list := []my_utils.IndexGraphEdge{}
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 1, Weight: 5})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 2, Weight: 7})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 0, End: 3, Weight: 3})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 4, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 1, End: 5, Weight: 10})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 2, End: 6, Weight: 1})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 3, End: 7, Weight: 11})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 7, Weight: 9})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 4, End: 1, Weight: 2})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 5, End: 7, Weight: 4})
	edge_list = append(edge_list, my_utils.IndexGraphEdge{Start: 6, End: 7, Weight: 6})

	adjacent_matrix_directed_graph_with_weights := RepresentDirectedGraphAsWeightedAdjacentMatrix(edge_list, 8)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[0][1], 5)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[0][2], 7)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[0][3], 3)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[1][4], 2)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[1][5], 10)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[2][6], 1)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[3][7], 11)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[4][7], 9)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[4][1], 2)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[5][7], 4)
	assert.Equal(t, adjacent_matrix_directed_graph_with_weights[6][7], 6)

}
