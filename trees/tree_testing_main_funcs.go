package main

import "awesomeProject1/trees/tree_traversals"

/*
	testing the predecessor and successor function
 */

//func main()  {
//	arr := []int{10, 5, 1, 7, 20, 40, 0}
//	root := tree_traversals.Constructing_bst(arr)
//
//	tree_traversals.Inorder_predecssor_successor_tester(root, arr)
//}


/*
	testing Left view traversal of BST
 */

func main(){
	arr := []int{10, 5, 1, 7, 20, 40, 0,8, 9,6,4,3,2, 60, 90, 100, 110 }
	root := tree_traversals.Constructing_bst(arr)

	tree_traversals.LeftViewOfBinarySearchTree(root)
}