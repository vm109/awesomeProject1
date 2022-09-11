package main

func RotateArray(arr [][]int, degree int) [][]int {
	number_of_times := degree / 90

	if number_of_times <= 0 {
		return arr
	} else {
		// first element will be pivot
		lengthOfRow := len(arr[0])
		lengthOfColumns := len(arr)
		// creating a new array
		newarr := make([][]int, 0, lengthOfColumns)
		for i := 0; i < lengthOfColumns; i++ {
			row := make([]int, lengthOfRow)
			newarr = append(newarr, row)
		}

		for i := 0; i < lengthOfRow; i++ {
			for j := 0; j < lengthOfColumns; j++ {
				newarr[j][i] = arr[i][lengthOfColumns-1-j]
			}
		}
		degree = degree - 90
		return RotateArray(newarr, degree)
	}
}
