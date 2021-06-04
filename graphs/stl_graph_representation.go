package graphs
/*
Graphs can be represented as adjacency matrix
or  adjacency lists
or we can use stl [ that is using a map like struct ]
 */


func Represent_Graph_Using_Map(edges []*GraphEdge) map[int][]int{
	graph_map := make(map[int][]int)
	for _, edge := range edges {
			graph_map[edge.StartNode] = append(graph_map[edge.StartNode],edge.EndNode)
			graph_map[edge.EndNode] = append(graph_map[edge.EndNode],edge.StartNode)
	}
	return graph_map
}