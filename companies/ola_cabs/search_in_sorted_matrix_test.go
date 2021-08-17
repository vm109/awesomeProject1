package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayOfIntegerArrays_FindANumberInMatrixArray(t *testing.T) {
	arrOfArr := ArrayOfIntegerArrays{
		ArrOfArr: [][]int{{3,30,38},{44,52,54},{57,60,69}},
	}
	found := arrOfArr.FindANumberInMatrixArray(3)
	assert.Equal(t, 1, found)
	found = arrOfArr.FindANumberInMatrixArray(62)
	assert.Equal(t, 0, found)
	found = arrOfArr.FindANumberInMatrixArray(69)
	assert.Equal(t, 1, found)
	found = arrOfArr.FindANumberInMatrixArray(65)
	assert.Equal(t, 0, found)
	found = arrOfArr.FindANumberInMatrixArray(44)
	assert.Equal(t, 1, found)
	found = arrOfArr.FindANumberInMatrixArray(35)
	assert.Equal(t, 0, found)
}
