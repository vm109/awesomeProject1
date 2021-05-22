package google

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTripletsInArrayWhoseSumIsZeroWithNoPossibleValues(t *testing.T) {
	arr := []int{1,2,3,4,5}
	triplet_arr := FindTripletsInArrayWhoseSumIsZero(arr)
	assert.Equal(t, 0, len(triplet_arr))
}

/*
Here the possible solutions are
{-1,0,1}
{-2,0,2}
 */
func TestFindTripletsInArrayWhoseSumIsZeroWithAPossibleSolution(t *testing.T) {
	arr := []int{-1,0,1,2,3,-2}
	triplet_map := FindTripletsInArrayWhoseSumIsZero(arr)
	fmt.Println(triplet_map)
	assert.Equal(t, 16, len(triplet_map))
}

func TestFindTripletsInArrayWhoseSumIsZeroWithTwoPossibleSolutions(t *testing.T) {
	arr := []int{0, -1, 2, -3, 1}
	triplet_map := FindTripletsInArrayWhoseSumIsZero(arr)
	fmt.Println(triplet_map)
	assert.Equal(t, 11, len(triplet_map))
}
