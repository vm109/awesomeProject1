package tree_traversals

import (
	"awesomeProject1/trees/tree_structs"
)
/*
Inorder traversal will always result in sorted order
 */
func traverse(node *tree_structs.Node, arr []int)[]int{
	if(node.Left != nil){
		arr = traverse(node.Left, arr)
	}
	arr = append(arr, node.Value)
	if(node.Right != nil){
		arr =traverse(node.Right, arr)
	}
	return arr
}

func InorderTraversal(root *tree_structs.Node)[]int{
	arr := []int{}
	return traverse(root, arr)
}