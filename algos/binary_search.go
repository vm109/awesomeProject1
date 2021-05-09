package algos

/*
Binary search is for searching an array by recursively
dividing the array to find the number.

Binary search can be done on a sorted array
 */

func Binary_search(sorted_arr []int, value, l, r int)int{
	mid := (l+r)/2
	if(sorted_arr[mid] < value){
		return Binary_search(sorted_arr, value, mid+1, r)
	} else if(sorted_arr[mid] > value){
		return Binary_search(sorted_arr, value, l, mid-1)
	} else if(sorted_arr[mid] == value){
		return  mid
	}else{
		return -1
	}
}