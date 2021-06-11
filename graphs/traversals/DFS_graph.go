package traversals

func dfs_visitor (graph map[int][]int, adjacent_verticies []int, visited_nodes map[int]bool) map[int]bool{
	for _, vertex := range adjacent_verticies {
		if(!visited_nodes[vertex]){
			visited_nodes[vertex] = true
			dfs_visitor(graph, graph[vertex], visited_nodes)
		}
	}
	return visited_nodes
}
func DFS_Graph(graph map[int][]int) map[int]bool{
	nodes := make(map[int]bool,0)
	for _, val := range graph{
		dfs_visitor(graph, val, nodes )
	}
	return nodes
}
