package traversals

import (
	"awesomeProject1/graphs"
)

type GraphRepresentedAsAdjacencyList struct {
	list map[int][]graphs.GraphNode
}

/*
Graph - Depth First Search :
Depth First Traversal is going deep into a given node when there are edges from a given node.
i.e, say there is edge from 2 -> 3 and 3 -> 4 and 4 -> 5 and 2->0

now the traversal is 2, 3, 4, 5 and then we will visit 0 [ since 0 is at level one but a different edge ]

so in essence
we traverse deeply of one side of a node and then go to the next edge and repeat it recursively until there
are no more non-visited nodes.
 */
func (adj_list GraphRepresentedAsAdjacencyList) DFS_Graph_with_adjacency_list(starting_key int, visited_nodes map[int]bool) []int{
	visited_nodes_in_order := make([]int,0)
	if(!visited_nodes[starting_key]){
		visited_nodes[starting_key] = true
		visited_nodes_in_order = append(visited_nodes_in_order, starting_key)
		for _, adjacent_node := range adj_list.list[starting_key] {
			visited_nodes_in_order = append(visited_nodes_in_order, adj_list.DFS_Graph_with_adjacency_list(adjacent_node.Data, visited_nodes)...)
		}
	}
	return visited_nodes_in_order
}
