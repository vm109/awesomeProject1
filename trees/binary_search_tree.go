/*
Binary search tree can consist of at most 2 nodes.
i.e; it can have 0,1 or 2 nodes

Left Subtree will have nodes always less than the parent
Right Subtree will have nodes always greater than parent
 */

package main

import "fmt"
import treestruct "awesomeProject1/trees/tree_structs"

func traverse_to_next_node(node *treestruct.Node, value int)  {
	if(node.Value < value) {
		if (node.Right != nil) {
			traverse_to_next_node(node.Right, value)
		} else {
			node.Right = &treestruct.Node{Value: value}
		}
	} else if(node.Value > value){
		if(node.Left != nil) {
			traverse_to_next_node(node.Left, value)
		}else{
			node.Left = &treestruct.Node{ Value: value}
		}
	}
}

func createBST(arr []int) *treestruct.Node{
	root := &treestruct.Node{ Value: arr[0]}
	for _,val := range arr[1:] {
		traverse_to_next_node(root, val)
	}
	return root
}

func traverse_a_bst(root *treestruct.Node){
	fmt.Println(root.Value)
	if(root.Left != nil){
		traverse_a_bst( root.Left)
	}
	if(root.Right != nil){
		traverse_a_bst(root.Right)
	}
}

//func  main()  {
//	arr:= []int{6,9,3,5,1,7,8,10}
//
//	traverse_a_bst(createBST(arr))
//}
