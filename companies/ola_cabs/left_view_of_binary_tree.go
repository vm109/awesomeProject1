package ola_cabs

/*
Problem Statement:
Given a Binary Tree, print Left view of it.
Left view of a Binary Tree is set of nodes visible when tree is visited from Left side.
The task is to complete the function leftView(), which accepts root of the tree as argument.
Example:
	Tree:

          1
       /     \
     2        3
   /   \    /   \
  4     5   6    7
   \
     8

In The above tree left view of the tree is 1, 2, 4, 8
*/

func (t Tree) LeftViewOfTree(arr []int) []int {
	if t.Left != nil {
		// go to left sub-tree
		arr = append([]int{t.Node},t.Left.LeftViewOfTree(arr)...)
		return arr
	} else if t.Right != nil {
		// go to right sub-tree
		arr = append([]int{t.Node},t.Right.LeftViewOfTree(arr)...)
		return arr
	} else {
		arr = append(arr, t.Node)
		return arr
	}
}
