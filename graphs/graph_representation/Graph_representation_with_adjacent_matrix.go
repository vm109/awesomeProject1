package graph_representation

import (
	"github.com/myawesomeproject1/my_utils"
)

func RepresentGraphAsMatrixWhenEdgeListIsGiven(edge_list []*my_utils.IndexGraphEdge, vertices int32) [][]int {
	arr := make([][]int, vertices)
	for i := 0; i < int(vertices); i++ {
		arr[i] = make([]int, vertices)
	}
	for _, value := range edge_list {
		arr[value.Start][value.End] = 1
		arr[value.End][value.Start] = 1
	}
	return arr
}

/*
A weighted Adjacent Matrix is representing a graph as a matrix.
And each matrix cell which has value more than 0 represents the wight of that edge.
*/
func RepresentDirectedGraphAsWeightedAdjacentMatrix(edge_list []my_utils.IndexGraphEdge, vertices_length int32) [][]int {
	arr := make([][]int, vertices_length)
	for i := 0; i < int(vertices_length); i++ {
		arr[i] = make([]int, vertices_length)
	}
	for _, value := range edge_list {
		arr[value.Start][value.End] = int(value.Weight)
	}
	return arr
}

/*
Adjacent Matrix representation Of Graph
In this implementation Graph Edges are not passed.
we pass only the vertices.
So the Initial state of the graph starts with all zeros.
And the final state of the graph holds the values for all edges.
The following is a AdjacentMatrix representation for non-weighted graph
*/

type AdjacentMatrixRep [][]int16

func NewAdjacentMatrix(numberOfVertices int) AdjacentMatrixRep {
	var a AdjacentMatrixRep = make([][]int16, numberOfVertices, numberOfVertices)
	return a
}

func (a AdjacentMatrixRep) AddEdge(startEdge, endEdge int16) (err error) {
	if (startEdge >= (int16)(cap(a))) || (endEdge >= (int16)(cap(a))) {

		panic("vertices are out of bound of the adjacent matrix created. Create a new matrix of bigger size")
	}

	if cap(a[startEdge]) == 0 {
		a[startEdge] = make([]int16, cap(a), cap(a))
		a[startEdge][endEdge] = 1
	} else {
		a[startEdge][endEdge] = 1
	}
	if cap(a[endEdge]) == 0 {
		a[endEdge] = make([]int16, cap(a), cap(a))
		a[endEdge][startEdge] = 1
	} else {
		a[endEdge][startEdge] = 1
	}

	return err
}

func (a AdjacentMatrixRep) RemoveEdge(startEdge, endEdge int16) (err error) {
	if (startEdge >= (int16)(cap(a))) || (endEdge >= (int16)(cap(a))) {
		panic("vertices are out of bound of the adjacent matrix created. not allowed to remove")
	}

	if cap(a[startEdge]) == 0 {
		a[startEdge] = make([]int16, 0, cap(a))
		a[startEdge][endEdge] = 0
	} else {
		a[startEdge][endEdge] = 0
	}
	if cap(a[endEdge]) == 0 {
		a[endEdge] = make([]int16, 0, cap(a))
		a[endEdge][startEdge] = 0
	} else {
		a[endEdge][startEdge] = 0
	}

	return err
}

func (a AdjacentMatrixRep) GetAdjacentMatrixRepresentation() (matrix [][]int16) {
	return a
}
