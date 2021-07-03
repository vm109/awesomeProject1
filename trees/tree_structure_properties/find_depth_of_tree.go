package main

import (
	"awesomeProject1/trees/tree_structs"
	"fmt"
	"math"
)

// lets traverse the tree in-order [ that is left node first and root and then right node ]
func measure_node_depths(node *tree_structs.Node, depth, min_depth int) int{
	if(node.Left != nil){
		min_depth = measure_node_depths(node.Left, depth+1, min_depth)
	}
	if(node.Right != nil){
		min_depth = measure_node_depths(node.Right, depth+1, min_depth)
	}
	if(node.Right == nil && node.Left == nil){
		fmt.Println("leaf node ", node.Value)
		if(depth< min_depth){
			min_depth = depth
		}
	}
	return min_depth
}

/*
To find the max depth we recursively traverse across nodes and if we find a lef node see if the
depth is more than max_depth replace the max_depth
 */
func measureMaxDepth(tree *tree_structs.Node, depth, max_depth int) int{
	if(tree.Left != nil){
		max_depth = measureMaxDepth(tree.Left, depth+1, max_depth)
	}
	if(tree.Right != nil){
		max_depth = measureMaxDepth(tree.Right, depth+1, max_depth)
	}
	if(tree.Right == nil && tree.Left == nil){
		fmt.Println("leaf node ", tree.Value)
		if(depth > max_depth){
			max_depth = depth
		}
	}
	return max_depth
}
// Depth or Height of a tree is the number of levels a node is from the root.
// A level in tree is parent_node -> child_node.
func FindMinimumDepthOfTree(tree *tree_structs.Node) (min_depth int){
	min_depth = measure_node_depths(tree, 0, math.MaxInt64)
	return
}

func FindMaxDepthOfTree(tree *tree_structs.Node)(max_depth int){
	max_depth = measureMaxDepth(tree, 0, -1)
	return
}


