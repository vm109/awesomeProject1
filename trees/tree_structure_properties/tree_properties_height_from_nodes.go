package main

import (
	"math"
)

func minimum_height_of_tree_from_nodes(nodes int)int{
	return int(math.Log2(float64(nodes+1))-1)
}

func max_height_of_tree_from_number_of_nodes(){

}


//func main(){
//	// minimum height of 1 nodes should be 0
//	fmt.Println("height of one node",minimum_height_of_tree_from_nodes(15))
//}