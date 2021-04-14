package main

/*
These are simple functions to calculate the max number of nodes when
height of the binary tree is given

Binary Tree:
A binary tree has max of 2 nodes,
i.e it can have 0,1 or 2 children for a given node
 */

import (
	tree "awesomeProject1/trees/tree_structs"
	"math"
)

type Tree  tree.Tree
func (t *Tree) calculate_tree_nodes() int{
 return int(math.Pow(2, float64(t.Height +1)))-1
}
//
//func main()  {
//	t := &Tree{ Height: 3}
//	nodes := t.calculate_tree_nodes()
//	fmt.Println(nodes)
//}