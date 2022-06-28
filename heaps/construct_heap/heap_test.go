package construct_heap

import (
	"fmt"
	"testing"
)

func TestConstructHeap(t *testing.T) {
	arr := []int{4, 10, 3, 5, 1}
	arr = ConstructHeap(arr)
	fmt.Println(arr)
}

func TestConstructHeap1(t *testing.T) {
	arr := []int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17}
	arr = ConstructHeap(arr)
	fmt.Println(arr)
}
