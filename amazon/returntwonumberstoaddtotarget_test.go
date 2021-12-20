package amazon

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReturnTwoNumbersToAddValueToTarget(t *testing.T) {
	arr := []int{1, 4, 3, 5, 2, 6, 9}
	target := 3

	pair := ReturnTwoNumbersToAddValueToTarget(arr, target)
	require.Equal(t, pair, []int{1, 2})
}

func TestReturnTwoNumbersFromTheArrayToAddToTarget_HashSolution(t *testing.T) {
	arr := []int{1, 4, 3, 5, 2, 6, 9}
	target := 3

	pair := ReturnTwoNumbersFromTheArrayToAddToTarget_HashSolution(arr, target)
	require.Equal(t, pair, []int{1, 2})
}