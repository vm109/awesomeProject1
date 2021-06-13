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
