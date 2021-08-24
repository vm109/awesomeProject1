package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
      3
    /   \
   1     2
 */
func TestTree_ConnectAllNodesAtSameLevelOfABinaryTree_3Node_Tree(t *testing.T) {
	tree := &Tree{Node: 3,}
	tree.Left = &Tree{Node: 1,}
	tree.Right = &Tree{Node: 2,}

	connected_tree := tree.ConnectAllNodesAtSameLevelOfABinaryTree()
	assertion_values := []int{3, 1, 2}
	i := 0
	for connected_tree != nil {
		assert.Equal(t, assertion_values[i], connected_tree.Node)
		connected_tree = connected_tree.Right
		i++
	}
}

/*
Input:
      10
    /   \
   20   30
  /  \
 40  60

output: 10 -> 20 -> 30 -> 40 -> 60
 */
func TestTree_ConnectAllNodesAtSameLevelOfABinaryTree_5Node_Tree(t *testing.T) {
	tree := &Tree{Node: 10,}
	tree.Left = &Tree{Node: 20,}
	tree.Right = &Tree{Node: 30,}
	pointer := tree.Left
	pointer.Left = &Tree{Node: 40,}
	pointer.Right = &Tree{Node: 60,}

	connected_tree := tree.ConnectAllNodesAtSameLevelOfABinaryTree()
	assertion_values := []int{10, 20, 30, 40, 60}
	i := 0
	for connected_tree != nil {
		assert.Equal(t, assertion_values[i], connected_tree.Node)
		connected_tree = connected_tree.Right
		i++
	}
}
