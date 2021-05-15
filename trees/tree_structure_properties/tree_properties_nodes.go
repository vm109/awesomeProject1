package main

import (
	"math"
)

func tree_number_of_nodes_at_a_level(level int) int{
	return int(math.Pow(2, float64(level)))
}

func total_number_of_nodes_by_level(level int) int{
	return int(math.Pow(2, float64(level+1)) - 1)
}

//func main(){
//	// find number of nodes in tree level
//	number_of_nodes:= tree_number_of_nodes_at_a_level(0)
//	fmt.Println("number of nodes in level 0 -", number_of_nodes)
//
//	number_of_nodes = tree_number_of_nodes_at_a_level(2)
//	fmt.Println("number of nodes in level 2 -", number_of_nodes)
//
//	// finding total nodes
//	toal_number_of_nodes_by_level := total_number_of_nodes_by_level(0)
//	fmt.Println("total_nodes tree of level 0 - ", toal_number_of_nodes_by_level)
//
//	toal_number_of_nodes_by_level = total_number_of_nodes_by_level(2)
//	fmt.Println("total_nodes tree of  level 2 - ",toal_number_of_nodes_by_level)
//}
