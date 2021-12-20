package walmart_labs

import "fmt"

/*
Longest Consecutive Subsequence:
Problem: Find the Length of longest sub-sequence integers
Consecutive numbers can be in any order
*/

func LongestConsecutiveSubsequence(sequence []int) (longestSequence int) {
	longestSequence = 0
	a := quickPartition(sequence, 0, len(sequence)-1)
	fmt.Println(a)
	longestSequence = subSequence(a)

	return longestSequence
}

// partition
func quickPartition(arr []int, start, end int) []int {
	if end-start > 0 {
		j := quickSortElements(arr, start, end)
		quickPartition(arr, start, j)
		quickPartition(arr, j+1, end)
	}
	return arr
}

// sort
// 2,6,1,9,4,5,3
func quickSortElements(arr []int, startIndex, endIndex int) int {
	pivot := arr[startIndex]
	i := startIndex
	j := endIndex
	for i < j {
		for i < len(arr) && arr[i] <= pivot {
			i++
		}

		for j > 0 && arr[j] > pivot {
			j--
		}

		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	if i > j {
		arr[startIndex] = arr[j]
		arr[j] = pivot
	}
	return j
}

func subSequence(a []int) int {
	longestSequence := 0
	tempLength := 1
	for i, _ := range a {
		if i < len(a)-1 {
			if a[i+1]-a[i] == 1 {
				tempLength += 1
			} else {
				tempLength = 1
			}
			if tempLength > longestSequence {
				longestSequence = tempLength
			}
		}
	}
	return longestSequence
}
