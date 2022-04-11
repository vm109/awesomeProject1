package search

type Array struct {
	Arr []interface{}
}
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

func (a Array) BinarySearch(searchValue int,initial, final int) interface{}{
	mid := (initial+final)/2

	if (final-initial)+1 == 2 {
		if a.Arr[initial].(int) == searchValue {
			return initial
		}else if a.Arr[final].(int) == searchValue{
			return final
		}else{
			return -1
		}
	}else if a.Arr[mid].(int) > searchValue{
		return  a.BinarySearch(searchValue, 0, mid)
	}else if a.Arr[mid].(int) < searchValue{
		return a.BinarySearch(searchValue, mid, final)
	}else if a.Arr[mid].(int) == searchValue{
		return mid
	}else{
		return -1
	}
}