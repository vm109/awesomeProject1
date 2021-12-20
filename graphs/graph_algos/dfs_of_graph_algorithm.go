package graph_algos

import "fmt"

/*
DFS - depth first search of graph
1. Visit first node.
2. Then Visit adjacent nodes.
3. Visit a node in adjacent nodes of first node and re-visit nodes adjacent to that node.

[0] [2,3,4]
[2] [5,6]
[3] [7,8]
 */

func Dfs_Of_Graph(graph map[int][]int, firstNode int){
	adjacentNodes := graph[firstNode]
	fmt.Println(firstNode)
	visited := make(map[int]int, 0)
	visited[firstNode] = 1
	visitAdjacentNodes(graph, adjacentNodes, visited)
}

func visitAdjacentNodes(graph map[int][]int, nodes []int, visited map[int]int) {
	for _, val := range nodes{
		if visited[val] != 1 {
			fmt.Println(val)
			visited[val] = 1
			if graph[val] != nil {
				visitAdjacentNodes(graph, graph[val], visited)
			}
		}
	}
}