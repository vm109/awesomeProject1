package ola_cabs

import "math"

/*
Problem Statement:
Given an array arr of N integers.
Find the contiguous sub-array with maximum sum.
*/

/*
Base Case:
If an array has negative numbers then we will have the problem of chosing
Sub-Arrays whose sum is maximum.

If all the Integers are Positive - then the entire array is the answer
*/

func (a ArrayOfIntegers) KadanesAlgorithmFindingContiguousMaxSumSubArray() []int {
	return findMaxSubArrayForMixedPositiveAndNegativeIntegers(a.Arr)
	//is_entire_arr_positive_integers := true
	//is_entire_arr_negative_integers := true
	//largest_negative_number := math.MinInt64
	//for _, val := range a.Arr {
	//	if val < 0 {
	//		is_entire_arr_positive_integers = false
	//		if val > largest_negative_number {
	//			largest_negative_number = val
	//		}
	//	} else if val > 0 {
	//		is_entire_arr_negative_integers = false
	//	}
	//}
	//if is_entire_arr_positive_integers {
	//	return a.Arr
	//} else if is_entire_arr_negative_integers {
	//	return []int{largest_negative_number}
	//} else {
	//	return []int{}
	//}
}

/*
a1 = [-1, 0, 2, -7, 4]
[-1] i=0
[-1,0]
[-1,0,2]
[-1,0,2,-7]
[-1,0,2,-7,4]
[0] i=1
[0,2]
a2 = [-1, 0, -1, 0, -1]
 */
func findMaxSubArrayForMixedPositiveAndNegativeIntegers(arr []int) []int {
	sum := math.MinInt64
	sum_cal := 0
	sub_arr := make([]int,0,0)
	for i:=0; i< len(arr); i++ {
		sum_cal = 0
		for j:= i; j<len(arr); j++{
			sum_cal += arr[j]
			if( sum_cal > sum){
				sum = sum_cal
				sub_arr = arr[i:j+1]
			}
		}
	}
	return sub_arr
}
