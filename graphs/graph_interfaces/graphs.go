package graph_interfaces

type Graph interface {
	AddEdgeIntoAdjacentList(source, destination int16) (err error)
	GetAdjacencyList() (adjList [][]int16)
	FindGraphCycle() (hasCycle bool)
}
