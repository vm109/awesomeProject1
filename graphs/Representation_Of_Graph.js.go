package graphs

import (
	"fmt"
	"strconv"
)

type Graph_Node struct {
	value int
	next *Graph_Node
}
func Add_edge(adjacency_list map[int][]int, u,v int) map[int][]int{
	if(adjacency_list[u] == nil){
		adjacency_list[u] = append(adjacency_list[u], u)
	}
		adjacency_list[u] = append(adjacency_list[u], v)
	if(adjacency_list[v] == nil){
		adjacency_list[v] = append(adjacency_list[v], v)
	}
		adjacency_list[v] = append(adjacency_list[v], u)
	return adjacency_list
}

func Print(adjacency_list map[int][]int){
	print_string := ""
	for _, array := range adjacency_list{
		for _, value := range array{
			print_string = print_string+ "->"+ strconv.Itoa(value)
		}
		fmt.Println(print_string)
		print_string = ""
	}
}

//func Wtf_is_wrong(){
//	a := []int{}
//	a = append(a, 1)
//	a = append(a, 2)
//	a = append(a, 3)
//	a = append(a, 4)
//	a = append(a, 5)
//	fmt.Println(a)
//}