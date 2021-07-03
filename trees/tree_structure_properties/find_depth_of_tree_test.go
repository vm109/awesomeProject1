package main

import (
	"awesomeProject1/trees/tree_structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
			9
		1		  2
	3       4   5
					6
						7
							8
 */

func buildTree_One()(tree *tree_structs.Node){
	tree = &tree_structs.Node{
		Value: 9,
	}
	node1 := &tree_structs.Node{
		Value: 1,
	}
	node2 := &tree_structs.Node{
		Value: 2,
	}
	node3 := &tree_structs.Node{
		Value: 3,
	}
	node4 := &tree_structs.Node{
		Value: 4,
	}
	node5 := &tree_structs.Node{
		Value: 5,
	}
	node6 := &tree_structs.Node{
		Value: 6,
	}
	node7 := &tree_structs.Node{
		Value: 7,
	}
	node8 := &tree_structs.Node{
		Value: 8,
	}
	tree.Left = node1
	tree.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node5.Right = node6
	node6.Right = node7
	node7.Right = node8
	return
}

func TestFindMinimumDepthOfTree(t *testing.T) {
	tree := buildTree_One()
	min_depth := FindMinimumDepthOfTree(tree)
	assert.Equal(t, min_depth, 2)
}

func TestFindMaxDepthOfTree(t *testing.T) {
	tree := buildTree_One()
	max_depth := FindMaxDepthOfTree(tree)
	assert.Equal(t, max_depth, 5)
}

/*
			9
      1		       2
   3      4     5
8	  7				6
						10

*/

func buildTree_Two()(tree *tree_structs.Node){
	tree = &tree_structs.Node{
		Value: 9,
	}
	node1 := &tree_structs.Node{
		Value: 1,
	}
	node2 := &tree_structs.Node{
		Value: 2,
	}
	node3 := &tree_structs.Node{
		Value: 3,
	}
	node4 := &tree_structs.Node{
		Value: 4,
	}
	node5 := &tree_structs.Node{
		Value: 5,
	}
	node6 := &tree_structs.Node{
		Value: 6,
	}
	node7 := &tree_structs.Node{
		Value: 7,
	}
	node8 := &tree_structs.Node{
		Value: 8,
	}
	node10 := &tree_structs.Node{
		Value: 10,
	}
	tree.Left = node1
	tree.Right = node2
	node1.Left = node3
	node1.Right = node4
	node3.Left = node8
	node3.Right = node7
	node2.Left = node5
	node5.Right = node6
	node6.Right = node10
	return
}

func TestFindMinimumDepthOfTreeTwo(t *testing.T) {
	tree := buildTree_Two()
	min_depth := FindMinimumDepthOfTree(tree)
	assert.Equal(t, min_depth, 2)
}

func TestFindMaxDepthOfTreeTwo(t *testing.T) {
	tree := buildTree_Two()
	max_depth := FindMaxDepthOfTree(tree)
	assert.Equal(t, max_depth, 4)
}

func buildTree_Three()(tree *tree_structs.Node) {
	tree = &tree_structs.Node{
		Value: 9,
	}
	return
}

func TestFindMinimumDepthOfTree_Three(t *testing.T) {
	tree := buildTree_Three()
	min_depth := FindMinimumDepthOfTree(tree)
	assert.Equal(t, min_depth, 0)
}

func TestFindMaxDepthOfTree_Three(t *testing.T) {
	tree := buildTree_Three()
	max_depth := FindMaxDepthOfTree(tree)
	assert.Equal(t, max_depth, 0)
}