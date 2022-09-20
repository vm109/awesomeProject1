package sep_15_2022

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPairOfnumbersInAnArrayWhoseProductIsValueX(t *testing.T) {
	arr := []int{2, 3, 4, 5, 1, 6, 8, 2}
	data := Data{Arr: arr}

	data.FindPairsOfNumbersInAnArrayWhoseProductIsValueX(12)

	assert.Equal(t, 3, len(data.FinalPairs))
}

func TestFindPairOfnumbersInAnArrayWhoseProductIsValueXOnePairs(t *testing.T) {
	arr := []int{2, 3, 4, 5, 1, 6, 8, 2}
	data := Data{Arr: arr}

	data.FindPairsOfNumbersInAnArrayWhoseProductIsValueX(5)

	assert.Equal(t, 1, len(data.FinalPairs))
}

func TestFindPairOfnumbersInAnArrayWhoseProductIsValueXNoPairs(t *testing.T) {
	arr := []int{2, 3, 4, 5, 1, 6, 8, 2}
	data := Data{Arr: arr}

	data.FindPairsOfNumbersInAnArrayWhoseProductIsValueX(9)

	assert.Equal(t, 0, len(data.FinalPairs))
}
