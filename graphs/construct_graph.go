package graphs

import "fmt"

type GraphNode struct{
	Data int
}

type GraphEdge struct{
	StartNode int
	EndNode int
}

func Print_Graphs() {
	adjacency_matrix := [][]int{[]int{0, 1, 0}, []int{1, 0, 1}, []int{0, 0, 1}}

	for i, value := range adjacency_matrix{
		for j,jValue := range value {
			if(jValue == 1){
				fmt.Println("there is a graph edge from ", i, j)
			}
		}
	}
}


func CreateAdjacenyMatricFromEdges(edges []GraphEdge) map[int][]*GraphNode{
	graph_map := make(map[int][]*GraphNode, len(edges))
	for _, edge := range edges{
		node_start := &GraphNode{
			Data: edge.StartNode,
		}
		node_end := &GraphNode{
			Data : edge.EndNode,
		}
		graph_map[edge.StartNode] = append(graph_map[edge.StartNode], node_end)
		graph_map[edge.EndNode] = append(graph_map[edge.EndNode], node_start)
	}

	return graph_map
}

