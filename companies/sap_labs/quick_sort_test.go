package sap_labs

import (
	"fmt"
	"testing"
)

func TestQuickSortArray(t *testing.T) {
	arr := []int{3,1,4,10,0,9,6,-1}
	arr = QuickSortArray(arr,0, len(arr)-1)
	fmt.Println(arr)
}
