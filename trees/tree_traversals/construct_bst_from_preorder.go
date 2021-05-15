package tree_traversals

import "awesomeProject1/trees/tree_structs"
/*
Given a Preorder array
construct the BST
BST - binary search tree is where the left child is less than root and
the right child is greater than the root

Say for this array {10, 5, 1, 7, 40, 50}

Preorder - root, left, right

root - 10, left - 5, left -1, 7 - right of 5
 */

func construct(root *tree_structs.Node, value int){
	if(value < root.Value ){
		if( root.Left != nil){
			construct(root.Left, value)
		}else{
			root.Left = &tree_structs.Node{ Value: value}
		}
	}else if (value > root.Value){
		if( root.Right != nil){
			construct(root.Right, value)
		}else{
			root.Right = &tree_structs.Node{Value: value}
		}
	}
}

func Constructing_bst(inoder_arr []int) *tree_structs.Node{
	root := &tree_structs.Node{
		Value : inoder_arr[0],
	}
	for _,value := range inoder_arr[1:]{
		construct(root, value)
	}

	return root
}
