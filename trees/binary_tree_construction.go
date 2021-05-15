package main

import (
	"awesomeProject1/my_utils"
	"fmt"
)
func Construct_tree(arr []int){
	queue := my_utils.Queue{
		make([]interface{},0),
	}
	root := &Node{arr[0], nil, nil}
	queue.Enqueue(root)
	i:= 1
	for i+1 < len(arr) {
		pointer := (queue.Dequeue()).(*Node)
		left := &Node{arr[i], nil, nil}
		right := &Node{arr[i+1], nil, nil}
		pointer.left = left
		pointer.right = right
		queue.Enqueue(left)
		queue.Enqueue(right)
		i = i+2
	}
	fmt.Println(root)
}

//func main()  {
//	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14}
//	Construct_tree(arr)
//}