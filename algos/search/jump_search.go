package search

import "math"

func find_jump_step(arr_length int) int{
	return int(math.Sqrt(float64(arr_length)))
}

func step_back(arr []int,value, start_index_to_step_back, step int) int{
	i := start_index_to_step_back
	end := start_index_to_step_back - step
	index := 0
	for i >= end {
		if(arr[i] == value){
			index = i
			break
		}
		i = i-1
	}
	return index
}
func Jump_search(arr []int, value int) int{
	step := find_jump_step(len(arr))
	i:= 0
	index := -1
	for i < len(arr){
		if(arr[i] == value){
			index= i
		}else if( arr[i] > value) {
			 index = step_back(arr, value, i, step)
		}
		i = i + step
	}
	return index
}
