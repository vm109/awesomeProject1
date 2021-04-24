package amazon

import (
	"fmt"
	"math"
)

/*
Given  a array [10, 20, 30, 40]
there will be matrix
a = {10 x 20 }
b = {20 x 30 }
c = {30 x 40 }
(ab)c = (10x20x30)+(10x30x40)
a(bc) = (20x30x40)+(10x20x40)

array [10, 30, 40, 50 ,3]
a = {10x30}
b=  {30x40}
c = {40x50}
d = {50x3}

(a(b(cd))) = (40x50x3)+(30x40x3)+(10x30x3)

((a(bc))d) = (30x40x50)+(10x30x50)+(10x50x3)
 */

func find_the_lowest_number_and_index(arr []int) (int, int){
	lowest := math.MaxInt16
	var index int
	for i, value := range arr{
		if(value < lowest){
			lowest = value
			index = i
		}
	}
	return lowest, index
}
func copy_arr_without_element(arr []int, index int) []int{
	new_arr := []int{}
	for i, value := range arr {
		if (index != i){
			new_arr = append(new_arr, value)
		}
	}
	return new_arr
}
func condition_to_move_forward_or_backward(arr []int, index int) string{
	if(len(arr) == 3) {
		return "all"
	}else if(index-2 < 0) {
		return "right"
	} else if(index+2 > len(arr)-1){
		return "left"
	}else if(index-2 >=0 && index+1 < len(arr) && arr[index-1] > arr[index+1]){
		return "left"
	}else if(index-1 >= 0 && index+2 < len(arr) && arr[index+1] > arr[index-1]){
		return "right"
	}
	return "not defined"
}

func find_operations(arr []int, index int, direction string) int{
	total_operations := 1
	if(direction == "all"){
		for i := 0; i < len(arr); i++ {
			total_operations = total_operations * arr[i]
		}
	}else if( direction == "left"){
		for i:= index; i > index-3; i-- {
			total_operations = total_operations * arr[i]
		}
	} else if ( direction == "right"){
		for i:= index; i< index+3; i++{
			total_operations = total_operations * arr[i]
		}
	}
	return total_operations
}
func Print_bracket_for_matrix_multiplication(size_arr []int, total int){
	_, index := find_the_lowest_number_and_index(size_arr)
	total_operations := 1
	new_arr := []int{}
	if(condition_to_move_forward_or_backward(size_arr, index) == "all") {
		total_operations = find_operations(size_arr, index, "all")
		new_arr = copy_arr_without_element(size_arr, 1)
	}else if(condition_to_move_forward_or_backward(size_arr, index) == "left"){
		total_operations = find_operations(size_arr, index, "left")
		new_arr = copy_arr_without_element(size_arr, index-1)
	}else if( condition_to_move_forward_or_backward(size_arr, index) == "right"){
		total_operations = find_operations(size_arr, index, "right")
		new_arr = copy_arr_without_element(size_arr, index+1)
	}
	fmt.Println(total_operations)
	fmt.Println(new_arr)
	total = total + total_operations
	fmt.Println(total)
	if(len(new_arr) > 2) {
		Print_bracket_for_matrix_multiplication(new_arr, total)
	}
}

