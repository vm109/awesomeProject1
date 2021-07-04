package graph_representation

import "awesomeProject1/graphs"

/*
Adjacent Nodes: The nodes which are connected to a given node are called adjacent nodes
Adjacency List: The list of all nodes with their corresponding adjacent nodes is called adjacent list

What is adjacency list is used for?
Adjacency list is used to represent graphs.

What is the input for graph representation?
Grpahs representation input can be the list of edges. An Edge is the connection between Node A and Node B
*/

func RepresentGraphAsAdjacencyListWhenEdgesListIsGiven(graph_edges []graphs.GraphEdge) map[int][]graphs.GraphNode{
	graph_adjacency_list := make(map[int][]graphs.GraphNode, 0)
	for _, graph_edge := range graph_edges {
		if(graph_adjacency_list[graph_edge.StartNode] == nil) {
			newarrStart := make([]graphs.GraphNode,0)
			newarrStart = append(newarrStart, graphs.GraphNode{ Data: graph_edge.EndNode})
			graph_adjacency_list[graph_edge.StartNode] = newarrStart
		}else {
			graph_adjacency_list[graph_edge.StartNode] = append(graph_adjacency_list[graph_edge.StartNode], graphs.GraphNode{ Data: graph_edge.EndNode})
		}
	}
	return graph_adjacency_list
}
