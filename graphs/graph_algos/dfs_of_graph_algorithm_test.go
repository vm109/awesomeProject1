package graph_algos

import "testing"

func TestDfs_Of_Graph(t *testing.T) {
	graph := make(map[int][]int, 0)
	graph[0] = []int{1,3}
	graph[1] = []int{2}
	graph[2] = []int{0,3}
	graph[3] = []int{3,4}
	Dfs_Of_Graph(graph, 0)
}

func TestDfs_Of_Graph2(t *testing.T) {
	graph := make(map[int][]int, 0)
	graph[0] = []int{1,9}
	graph[1] = []int{2}
	graph[2] = []int{0,3}
	graph[9] = []int{3}
	Dfs_Of_Graph(graph, 0)
}
