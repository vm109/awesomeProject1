package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayOfIntegers_FindMissingNumberInGivenArray(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{ 1,3,4,5},
	}
	missing_number := arr.FindMissingNumberInGivenArray()
	assert.Equal(t, 2, missing_number)
}

func TestArrayOfIntegers_FindMissingNumberInGivenArray2(t *testing.T) {
	arr := ArrayOfIntegers{
		Arr: []int{1,2,3,4,6,7,8},
	}
	missing_number := arr.FindMissingNumberInGivenArray()
	assert.Equal(t, 5, missing_number)
}
