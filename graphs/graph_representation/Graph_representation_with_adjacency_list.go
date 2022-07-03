package graph_representation

import (
	"github.com/myawesomeproject1/graphs"
	"github.com/pkg/errors"
)

/*
Adjacent Nodes: The nodes which are connected to a given node are called adjacent nodes
Adjacency List: The list of all nodes with their corresponding adjacent nodes is called adjacent list

What is adjacency list is used for?
Adjacency list is used to represent graphs.

What is the input for graph representation?
Grpahs representation input can be the list of edges. An Edge is the connection between Node A and Node B
*/

func RepresentGraphAsAdjacencyListWhenEdgesListIsGiven(graph_edges []graphs.GraphEdge) map[int][]graphs.GraphNode {
	graph_adjacency_list := make(map[int][]graphs.GraphNode, 0)
	for _, graph_edge := range graph_edges {
		if graph_adjacency_list[graph_edge.StartNode] == nil {
			newarrStart := make([]graphs.GraphNode, 0)
			newarrStart = append(newarrStart, graphs.GraphNode{Data: graph_edge.EndNode})
			graph_adjacency_list[graph_edge.StartNode] = newarrStart
		} else {
			graph_adjacency_list[graph_edge.StartNode] = append(graph_adjacency_list[graph_edge.StartNode], graphs.GraphNode{Data: graph_edge.EndNode})
		}
	}
	return graph_adjacency_list
}

/*
Adjacency List: Implementation
a. Create Adjacency List Array
b. Add Adjacent Nodes
c. Remove Adjacent Nodes
d. Return Adjacent List
*/

type AdjacentListRep [][]int16

func NewAdjacencyListRep(numberOfVertices int16) AdjacentListRep {
	var adjListRep AdjacentListRep = make([][]int16, numberOfVertices, numberOfVertices)
	return adjListRep
}

func (a AdjacentListRep) AddEdgeIntoAdjacentList(source, destination int16) (err error) {
	if (source >= (int16)(cap(a))) || (destination >= (int16)(cap(a))) {
		return errors.New("source or destination is greater than the allowed capacity of the adjacency list. Please create a new adjacency list and try")
	}

	if len(a[source]) == 0 {
		list := make([]int16, 0, cap(a))
		list = append(list, destination)
		a[source] = list
	} else {
		a[source] = append(a[source], destination)
	}

	if len(a[destination]) == 0 {
		list := make([]int16, 0, cap(a))
		list = append(list, source)
		a[destination] = list
	} else {
		a[destination] = append(a[destination], source)
	}
	return
}

func (a AdjacentListRep) GetAdjacencyList() (adjList [][]int16) {
	return a
}
