package graph_representation

import (
	"awesomeProject1/my_utils"
	"fmt"
)

func PrintEdgesAndTheWeight(edge_list []*my_utils.GraphEdge) {
	for _, value := range edge_list{
		fmt.Println(value.Start, "-", value.End, " weight", value.Weight)
	}
}

func GetVerticesFromEdgeList(edge_list []*my_utils.GraphEdge)map[string]int{
	vertices := make(map[string]int,0)
	for _, value := range edge_list {
		if(vertices[value.Start]==0){
			vertices[value.Start] = 1
		}
		if(vertices[value.End] == 0){
			vertices[value.End] = 1
		}
	}
	return  vertices
}
