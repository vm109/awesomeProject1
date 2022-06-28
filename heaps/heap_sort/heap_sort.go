package heap_sort

import (
	"fmt"
	"heaps/construct_heap"
)

func HeapSortArray(arr []int) (sortedArr []int) {
	heapified := construct_heap.ConstructHeap(arr)
	fmt.Println(heapified)
	arr = removeAndHeapify(heapified)
	fmt.Println(arr)
	return heapified
}

func removeAndHeapify(arr []int) (sortedArr []int) {
	if len(arr) > 2 {
		if sortedArr == nil {
			sortedArr = make([]int, 0)
		}
		sortedArr = append(sortedArr, arr[0])
		arr[0] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		sortedArr = append(sortedArr, removeAndHeapify(construct_heap.ConstructHeap(arr))...)
	} else if len(arr) == 2 {
		if sortedArr == nil {
			sortedArr = make([]int, 0)
		}
		sortedArr = append(sortedArr, arr...)
	}
	return
}
