package sap_labs

/*
Quick Sort Problem Description:
Quick sort is divide and conquer algorithm.
Pick an element as pivot and partition given array around the pivot element
and quick sort elements in that partition.
*/

func QuickSortArray(arr []int, start, end int) []int {
	sortElements(arr, start, end)
	return arr
}

func sortElements(arry_to_sort []int, start, end int) []int {
	if end-start > 0 {
		j := partitionAlgorithm(arry_to_sort, start, end)
		sortElements(arry_to_sort, start, j)
		sortElements(arry_to_sort, j+1, end)
	}
	return arry_to_sort
}

func partitionAlgorithm(arr_to_sort []int, start, end int) int {
	pivot := arr_to_sort[start]
	i := start
	j := end
	for i < j {
		for i<len(arr_to_sort) && arr_to_sort[i] <= pivot {
			i++
		}
		for j >0 && arr_to_sort[j] > pivot {
			j--
		}
		if i < j {
			temp := arr_to_sort[i]
			arr_to_sort[i] = arr_to_sort[j]
			arr_to_sort[j] = temp
		}
	}
	if j < i {
		temp := arr_to_sort[start]
		arr_to_sort[start] = arr_to_sort[j]
		arr_to_sort[j] = temp
	}
	return j
}
