package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_BubbleSortInteger(t *testing.T) {
	arr := make([]interface{}, 0)
	arr = append(arr, 4, 5, 2, 1, 6, 0, 9)
	a := Array{
		Arr : arr,
	}

	sortedArray := a.BubbleSortInteger()
	fmt.Print(sortedArray)
}


func TestArray_InsertionSort(t *testing.T) {
	arr := make([]interface{}, 0)
	arr = append(arr, 4, 5, 2, 1, 6, 0, 9)
	a := Array{
		Arr : arr,
	}
	sortedArray := a.InsertionSort()
	assert.Equal(t, sortedArray, []interface{}{0,1,2,4,5,6,9})
}

func TestArray_MergeSort(t *testing.T) {
	arr := make([]interface{}, 0)
	arr = append(arr, 4, 5, 2, 1, 6, 0, 9)
	a := Array{
		Arr : arr,
	}
	sortedArray := a.MergeSort()
	assert.Equal(t, sortedArray, []interface{}{0,1,2,4,5,6,9})
}