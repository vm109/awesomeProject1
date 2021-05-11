package main

import (
	sort_algos "awesomeProject1/algos/sort"
	"fmt"
)
/*
Testing Bubble sort
 */
//func main(){
//	arr := []int{5,4,1,10,9,3,13,20,2}
//	sorted_arr := sort_algos.Bubble_sort(arr)
//	fmt.Println(sorted_arr)
//}

func main(){
	arr := []int{5,4,1,10,9,3,13,20,2}
	sorted_arr := sort_algos.Insertion_sort(arr)
	fmt.Println(sorted_arr)
}

