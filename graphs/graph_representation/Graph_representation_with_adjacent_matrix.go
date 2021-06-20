package graph_representation

import "awesomeProject1/my_utils"

func RepresentGraphAsMatrixWhenEdgeListIsGiven(edge_list []*my_utils.IndexGraphEdge, vertices int32)[][]int{
	arr := make([][]int, vertices)
	for i:= 0; i< int(vertices); i++ {
		arr[i] = make([]int, vertices)
	}
	for _, value := range edge_list{
		arr[value.Start][value.End] = 1
		arr[value.End][value.Start] =1
	}
	return arr
}

/*
A weighted Adjacent Matrix is representing a graph as a matrix.
And each matrix cell which has value more than 0 represents the wight of that edge.
 */
func RepresentDirectedGraphAsWeightedAdjacentMatrix(edge_list []my_utils.IndexGraphEdge, vertices_length int32)[][]int{
	arr := make([][]int, vertices_length)
	for i:= 0; i< int(vertices_length); i++ {
		arr[i] = make([]int, vertices_length)
	}
	for _, value := range edge_list{
		arr[value.Start][value.End] = int(value.Weight)
	}
	return arr
}
