package construct_heap

// []int{4, 10, 3, 5, 1}
//       10
//      /  \
//     6    5
//   /  \
//  4    1

// Algorithm:
// find the last non-leaf node in the array.
// and heapify all the nodes prior to that node in the array.

func ConstructHeap(arr []int) []int {
	lastNonLeafIndex := findLastNonLeafNode(arr)
	for i := lastNonLeafIndex; i >= 0; i-- {
		arr = maxHeapify(arr, i)
	}
	return arr
}

func findLastNonLeafNode(arr []int) (lastNonLeafIndex int) {
	arrLength := len(arr)
	lastNonLeafIndex = ((arrLength - 1) - 1) / 2
	return
}

func maxHeapify(arr []int, heapifyValueIndex int) []int {

	leftChildInArrIndx := (2 * heapifyValueIndex) + 1
	rightChildInArrIndx := (2 * heapifyValueIndex) + 2

	if leftChildInArrIndx < len(arr) && arr[heapifyValueIndex] < arr[leftChildInArrIndx] {
		swp := arr[heapifyValueIndex]
		arr[heapifyValueIndex] = arr[leftChildInArrIndx]
		arr[leftChildInArrIndx] = swp
		arr = maxHeapify(arr, leftChildInArrIndx)
	}
	if rightChildInArrIndx < len(arr) && arr[heapifyValueIndex] < arr[rightChildInArrIndx] {
		swp := arr[heapifyValueIndex]
		arr[heapifyValueIndex] = arr[rightChildInArrIndx]
		arr[rightChildInArrIndx] = swp
		arr = maxHeapify(arr, rightChildInArrIndx)
	}
	return arr
}
