package main

import "testing"

func TestRotateArray(t *testing.T) {
	arr := make([][]int, 3)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 3)
	}

	arr[0][0] = 1
	arr[0][1] = 2
	arr[0][2] = 3

	arr[1][0] = 4
	arr[1][1] = 5
	arr[1][2] = 6

	arr[2][0] = 7
	arr[2][1] = 8
	arr[2][2] = 9

	RotateArray(arr, 270)
}
