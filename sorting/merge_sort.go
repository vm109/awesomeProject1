package main

import "fmt"

func merge_sort(arr []int, lb, ub int) []int{
	mid := (lb+ub)/2
	if(lb<ub){
		merge_sort(arr, lb, mid)
		merge_sort(arr, mid+1, ub)
		merged_array:= merge(arr, lb, mid, ub)
		copy(arr[lb:ub+1], merged_array)
	}
	return arr
}

func merge(arr []int, lb, mid, ub int) []int{
	left_arr:= arr[lb:mid+1]
	right_arr:= arr[mid+1:ub+1]
	l:=0
	r:=0
	var final_arr []int
	for(l< len(left_arr) || r< len(right_arr)){
		if(l < len(left_arr) && r< len(right_arr)) {
			if (left_arr[l] < right_arr[r]) {
				final_arr = append(final_arr, left_arr[l])
				l++
			} else if left_arr[l] > right_arr[r] {
				final_arr = append(final_arr, right_arr[r])
				r++
			}
		}else{
			if(l < len(left_arr)){
				final_arr = append(final_arr, left_arr[l])
				l++
			}
			if(r < len(right_arr)){
				final_arr = append(final_arr, right_arr[r])
				r++
			}
		}
	}
	return  final_arr
}

func main()  {
	array := []int{11,5,8,1,0,19,2,12,9}
	arr := merge_sort(array, 0, len(array)-1)
	fmt.Println(arr)
}