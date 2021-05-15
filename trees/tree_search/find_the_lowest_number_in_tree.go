package main

import (
	"awesomeProject1/trees/tree_structs"
	"awesomeProject1/trees/tree_traversals"
	"fmt"
)


func Find_the_lowest_number_in_tree(root *tree_structs.Node) int{
	return tree_traversals.InorderTraversal(root)[0]
}


func main(){
	inorder_arr := []int{10, 5, 1, 7, 20, 40, 0}
	root := tree_traversals.Constructing_bst(inorder_arr)
	lowest_number := Find_the_lowest_number_in_tree(root)
	fmt.Println("lowest number", lowest_number)
}
