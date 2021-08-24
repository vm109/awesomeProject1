package ola_cabs

import "fmt"

/*
Problem Statement:

Given a binary tree, connect the nodes that are at same level. You'll be given an addition nextRight pointer for the same.

Initially, all the nextRight pointers point to garbage values. Your function should set these pointers to point next
right for each node.

       10                         10 ------> NULL
      / \                       /      \
     3   5       =>           3 ------> 5 --------> NULL
    / \   \                  /  \        \
   4   1   2               4 --> 1 -----> 2 -------> NULL

 */

func (t *Tree) ConnectAllNodesAtSameLevelOfABinaryTree() *Tree{
	same_level_nodes_connected_tree := &Tree{}
	queue := &Queue{}
	queue.Enqueue(t)
	connect(same_level_nodes_connected_tree, queue)
	return same_level_nodes_connected_tree
}

func connect(same_level_nodes_connected_tree *Tree, queue *Queue) *Tree{
	if( !queue.IsEmpty() ){
		node := queue.Dequeue().(*Tree)
		if same_level_nodes_connected_tree != nil {
			same_level_nodes_connected_tree.Node = node.Node
		}else{
			same_level_nodes_connected_tree = &Tree{}
			same_level_nodes_connected_tree.Node = node.Node
		}
		fmt.Println(node.Node)
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
		same_level_nodes_connected_tree.Right = connect(same_level_nodes_connected_tree.Right, queue)
	}
	return same_level_nodes_connected_tree
}