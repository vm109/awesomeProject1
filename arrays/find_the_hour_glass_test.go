package arrays

import (
	"fmt"
	"testing"
)

func TestFindHourglassSum(t *testing.T) {
	arr := [][]int32{
		[]int32{0,-4,-6,0,-7,-6 },
		[]int32{-1,-2,-6,-8,-3,-1},
		[]int32{-8,-4,-2,-8,-8,-6},
		[]int32{-3,-1,-2,-5,-7,-4},
		[]int32{-3,-5,-3,-6,-6,-6},
		[]int32{-3,-6,0,-8,-6,-7},
	}
	sum := FindHourglassSum(arr)
	fmt.Println(sum)
}
