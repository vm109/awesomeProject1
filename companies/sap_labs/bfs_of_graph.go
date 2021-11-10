package sap_labs

import "fmt"

func Bfs_Traversal_Of_Graph(adjList map[int][]int, startNode int) {
	visitedNodes := make(map[int]int, 0)
	bfs_queue := make([]int, 0)

	bfs_queue = append(bfs_queue, startNode)

	for len(bfs_queue) > 0 {
		popped_element := bfs_queue[0]
		if len(bfs_queue)>1 {
			bfs_queue = bfs_queue[1:]
		}else {
			bfs_queue = make([]int, 0)
		}
		if visitedNodes[popped_element] != 1 {
			visitedNodes[popped_element] = 1
			bfs_queue = append(bfs_queue, adjList[popped_element]...)
			fmt.Println(popped_element)
		}
	}
}
