package graph_cycle

import (
	"reflect"
)

type AdjacentList [][]int16

func NewAdjacentList(numberOfVertices int) AdjacentList {
	var adjList AdjacentList = make([][]int16, numberOfVertices, numberOfVertices)
	return adjList
}

func (a AdjacentList) AddEdgeIntoAdjacentList(source, destination int16) (err error) {
	if (source >= (int16)(cap(a))) || (destination >= (int16)(cap(a))) {
		panic("source or destination is greater than the allowed capacity of the adjacency list. Please create a new adjacency list and try")
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

func (a AdjacentList) GetAdjacencyList() (adjList [][]int16) {
	return
}

func (a AdjacentList) FindGraphCycle() (hasCycle bool) {
	visitedVerticesMap := make(map[int16]bool, 0)

	startingVertex := 0

	return findGraphCycle(a, (int16)(startingVertex), visitedVerticesMap)
}

func findGraphCycle(a [][]int16, startingVertex int16, visitedVerticesMap map[int16]bool) (foundCycle bool) {
	adjVertices := a[startingVertex]

	if len(reflect.ValueOf(visitedVerticesMap).MapKeys()) != len(a) {
		for _, adjVertex := range adjVertices {
			if !visitedVerticesMap[adjVertex] {
				visitedVerticesMap[adjVertex] = true
				findGraphCycle(a, adjVertex, visitedVerticesMap)
			} else if visitedVerticesMap[adjVertex] && adjVertex != startingVertex {
				return true
			}
		}
	}

	return false
}
