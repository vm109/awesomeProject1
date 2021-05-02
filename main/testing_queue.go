package main

import (
	//utils "awesomeProject1/my_utils"
	//"fmt"
	bfs "awesomeProject1/trees/breadth_first_search"
)

func main(){
	//queue := &utils.Queue{
	//	make([]interface{},0),
	//}
	//queue.Enqueue(2)
	//queue.Enqueue(3)
	//queue.Enqueue(8)
	//queue.Enqueue(10)
	//for  range queue.Queue {
	//	fmt.Println(queue.Dequeue())
	//}
	root := &bfs.Node{
		2, nil, nil,
	}
	bfs.Add_nodes_tree(root, 3)
	bfs.Add_nodes_tree(root, 1)
	bfs.Add_nodes_tree(root, 0)
	bfs.Bfs_traverse(root)
}
