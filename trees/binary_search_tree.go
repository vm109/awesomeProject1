/*
Binary search tree can consist of at most 2 nodes.
i.e; it can have 0,1 or 2 nodes

Left Subtree will have nodes always less than the parent
Right Subtree will have nodes always greater than parent
 */

package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
}

func traverse_to_next_node(node *Node, value int)  {
	if(node.value < value) {
		if (node.right != nil) {
			traverse_to_next_node(node.right, value)
		} else {
			node.right = &Node{value: value}
		}
	} else if(node.value > value){
		if(node.left != nil) {
			traverse_to_next_node(node.left, value)
		}else{
			node.left = &Node{ value: value}
		}
	}
}

func createBST(arr []int) *Node{
	root := &Node{ value: arr[0]}
	for _,val := range arr[1:] {
		traverse_to_next_node(root, val)
	}
	return root
}

func traverse_a_bst(root *Node){
	fmt.Println(root.value)
	if(root.left != nil){
		traverse_a_bst( root.left)
	}
	if(root.right != nil){
		traverse_a_bst(root.right)
	}
}

func  main()  {
	arr:= []int{6,9,3,5,1,7,8,10}

	traverse_a_bst(createBST(arr))
}
