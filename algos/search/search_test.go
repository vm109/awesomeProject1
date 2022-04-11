package search

import (
	"fmt"
	"testing"
)

func TestArray_BinarySearch(t *testing.T) {
	arr := make([]interface{},0)
	arr = append(arr, []interface{}{2,4,5,7,19,23,130,233,240}...)
	a := Array{
		Arr: arr,
	}
	element := a.BinarySearch(130, 0, len(a.Arr)-1)
	fmt.Println(element)
}
