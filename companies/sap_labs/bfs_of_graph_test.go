package sap_labs

import "testing"

/*
Testing BFS traversal of Graph
Graph is represented as adjacent list.
That is list of all nodes of given node in graph

Example:
0 -> 1
0 -> 2
1 -> 2
2 -> 3

this can be represented in adjList as
[0][1,2]
[1][2]
[2][3,0]
[3][3]
*/

func TestBfs_Traversal_Of_Graph(t *testing.T) {
	adjList := make(map[int][]int, 0)
	adjList[0] = []int{1, 2}
	adjList[1] = []int{2}
	adjList[2] = []int{3,0}
	adjList[3] = []int{3}
	Bfs_Traversal_Of_Graph(adjList, 0)
}

/*
One more Graph Example

1 -> 2, 3
2 -> 4, 5
3 -> 5
4 -> 5, 6
 */

func TestBfs_Traversal_Of_Graph_One(t *testing.T) {
	adjList := make(map[int][]int, 0)
	adjList[1] = []int{2,3}
	adjList[2] = []int{4,5}
	adjList[3] = []int{5}
	adjList[4] = []int{5,6}
	Bfs_Traversal_Of_Graph(adjList, 1)
}