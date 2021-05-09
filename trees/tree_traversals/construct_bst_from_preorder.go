package tree_traversals

import "fmt"

/*
Given a Preorder array
construct the BST
BST - binary search tree is where the left child is less than root and
the right child is greater than the root

Say for this array {10, 5, 1, 7, 40, 50}

Preorder - root, left, right

root - 10, left - 5, left -1, 7 - right of 5
 */

type Node struct {
	value int
	left *Node
	right *Node
}

func construct(root *Node, value int){
	if(value < root.value ){
		if( root.left != nil){
			construct(root.left, value)
		}else{
			root.left = &Node{ value: value}
		}
	}else if (value > root.value){
		if( root.right != nil){
			construct(root.right, value)
		}else{
			root.right = &Node{value: value}
		}
	}
}

func Constructing_bst(inoder_arr []int){
	root := &Node{
		value : inoder_arr[0],
	}
	for _,value := range inoder_arr[1:]{
		construct(root, value)
	}

	fmt.Println("printing the tree", *root)
}
