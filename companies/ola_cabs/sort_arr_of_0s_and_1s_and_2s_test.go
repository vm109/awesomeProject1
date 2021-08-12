package ola_cabs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayOfIntegers_SortArrayOf0S1SAnd2S_Base_case(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{0, 2, 1, 2, 0},
	}
	sorted_arr := arr.SortArrayOf0S1SAnd2S()

	assert.Equal(t, sorted_arr, []int{0, 0, 1, 2, 2})
}

func TestArrayOfIntegers_SortArrayOf0S1SAnd2S_ArrayContainsOtherElementsLike3(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{0, 2, 1, 2, 0},
	}
	sorted_arr := arr.SortArrayOf0S1SAnd2S()

	fmt.Println(sorted_arr)
}
