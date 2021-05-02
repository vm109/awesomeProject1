package breadth_first_search

import (
	util "awesomeProject1/my_utils"
	"fmt"
)
type Node struct{
	Value int
	Left *Node
	Right *Node
}
func Add_nodes_tree(tree *Node,  value int){
	if(tree.Value > value ){
		if(tree.Left != nil ){
			Add_nodes_tree(tree.Left, value)
		}else{
			tree.Left = &Node{value, nil, nil }
		}
	}else if( tree.Value < value ){
		if(tree.Right != nil ){
			Add_nodes_tree(tree.Right, value)
		}else{
			tree.Right = &Node{value, nil, nil}
		}
	}
}

func traverse(queue *util.Queue){

	for  len(queue.Queue) > 0 {
		dequeued_item := (queue.Dequeue()).(Node)
		fmt.Println(dequeued_item.Value)
		if(dequeued_item.Left != nil) {
			queue.Enqueue(*dequeued_item.Left)
		}
		if(dequeued_item.Right != nil) {
			queue.Enqueue(*dequeued_item.Right)
		}
	}
}

func Bfs_traverse(tree *Node){
	queue := &util.Queue{
		 make([]interface{},0),
	}
	queue.Enqueue(*tree)
	traverse(queue)
}
