package ola_cabs

/*
Problem Statement:
Given a matrix mat[][] of size N x M, where every row and column is sorted in increasing order,
and a number X is given. The task is to find whether element X is present in the matrix or not.

Example 1:
Input -> mat[][] = 3 30 38
                   44 52 54
                   57 60 69
and find number '62'
Output: 0

Example 2:

mat[][] = 18 21 27 38 55 67
x = 55
output: 1
*/

func (a ArrayOfIntegerArrays) FindANumberInMatrixArray(number int) int {
	found := 0
	for _, arrOfIntegers := range a.ArrOfArr {
		if number >= arrOfIntegers[0] {
			found = findByBinarySearch(arrOfIntegers, number)
			if found == 1 {
				break
			}
		}
	}
	return found
}

func findByBinarySearch(arr []int, number int) int {
	mid := len(arr) / 2
	if mid > 0 && arr[mid] > number {
		return findByBinarySearch(arr[:mid], number)
	} else if mid > 0 && arr[mid] < number {
		return findByBinarySearch(arr[mid:], number)
	} else if arr[mid] == number {
		return 1
	} else {
		return 0
	}
}
