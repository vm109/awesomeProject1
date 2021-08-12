package ola_cabs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayOfIntegers_KadanesAlgorithmFindingContiguousMaxSumSubArray_AllPositiveNumbers(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{1,2,4,5,0,9},
	}
	sub_arr := arr.KadanesAlgorithmFindingContiguousMaxSumSubArray()
	fmt.Println(sub_arr)
	assert.Equal(t, sub_arr, arr.Arr)
}

func TestArrayOfIntegers_KadanesAlgorithmFindingContiguousMaxSumSubArray_AllNegativeNumbers(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{-2, -4, -3, -1, -5 },
	}
	sub_arr := arr.KadanesAlgorithmFindingContiguousMaxSumSubArray()
	fmt.Println(sub_arr)
	assert.Equal(t, sub_arr, []int{-1})
}

func TestArrayOfIntegers_KadanesAlgorithmFindingContiguousMaxSumSubArray_MixedIntegers(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{-1, 0, 2, -7, 4},
	}

	sub_arr := arr.KadanesAlgorithmFindingContiguousMaxSumSubArray()
	fmt.Println(sub_arr)
	assert.Equal(t, sub_arr, []int{4})
}

func TestArrayOfIntegers_KadanesAlgorithmFindingContiguousMaxSumSubArray_MixedIntegers_One(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{1,2,3,-2,5},
	}

	sub_arr := arr.KadanesAlgorithmFindingContiguousMaxSumSubArray()
	fmt.Println(sub_arr)
	assert.Equal(t, sub_arr, []int{1,2,3,-2,5})
}
