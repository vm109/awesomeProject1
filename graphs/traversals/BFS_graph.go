package traversals

import "fmt"

func BFS_graph(graph *map[int][]int) {
	visited_nodes := map[int]bool{}
	for _, adjacent_arr := range *graph {
		for _, value := range adjacent_arr {
			if(visited_nodes[value] == false) {
				fmt.Print(value," ")
				visited_nodes[value] = true
			}
		}
	}
}
