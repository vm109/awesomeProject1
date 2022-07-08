package graph_representation

import (
	"github.com/pkg/errors"
)

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
