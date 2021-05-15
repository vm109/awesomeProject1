package main

import (
	"fmt"
	"math"
)

/*
if number of leaf nodes = L
and number of internal nodes = I
L = I+1
 */
func leaf_nodes_and_internal_nodes(level int) bool{
	leaf_nodes := int(math.Pow(2,float64(level)))
	total_nodes := int(math.Pow(2, float64(level+1))-1)
	fmt.Println("leaf nodes", leaf_nodes)
	fmt.Println("total_nodes", total_nodes)
	internal_nodes := total_nodes - leaf_nodes
	fmt.Println(internal_nodes)
	return internal_nodes + 1 == leaf_nodes
}

func main(){
	leaf_nodes_and_internal_nodes(3)
}
