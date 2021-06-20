package graph_algos

import (
	"awesomeProject1/graphs/graph_representation"
	utils "awesomeProject1/my_utils"
	"math"
)

func find_next_vertex(directed_graph [][]int, source_vertex , distance int32, visited_nodes_arr []int32) (int32, []int32){
	min_value := math.MaxInt32
	min_vertex := -1
	for index, value := range directed_graph[source_vertex] {
		if(value>0 && value < min_value){
			min_value = value
			min_vertex = index
		}
	}
	if(min_vertex != -1) {
		distance += int32(min_value)
		visited_nodes_arr = append(visited_nodes_arr, int32(min_vertex))
		return find_next_vertex(directed_graph, int32(min_vertex), distance, visited_nodes_arr)
	}
	return int32(distance), visited_nodes_arr
}
func visit_nodes(directed_graph [][]int, source_vertex int32)(int32, []int32){
		total_distance := 0
		arr := []int32{ source_vertex }
		distance, visited_nodes_arr := find_next_vertex(directed_graph, source_vertex, int32(total_distance), arr)
		return distance, visited_nodes_arr
}

func Dijkstras_Find_Shortest_Path(weighted_graph []utils.IndexGraphEdge, vertices, source_vertex int32)(int32, []int32){
	weighted_directed_graph := graph_representation.RepresentDirectedGraphAsWeightedAdjacentMatrix(weighted_graph, vertices)
	return visit_nodes(weighted_directed_graph, source_vertex)
}
