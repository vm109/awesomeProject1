package heap_sort

import "testing"

func TestHeapSortArray(t *testing.T) {
	arr := []int{4, 3, 9, 2, 10, 15, 8}
	HeapSortArray(arr)
}

func TestHeapSortArray1(t *testing.T) {
	arr := []int{4, 3, 9, 2, 10, 15, 8, 99, 1, 13, 78, -1}
	HeapSortArray(arr)
}
