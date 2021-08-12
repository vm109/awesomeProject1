package ola_cabs

/*
Problem Statement:
Given an array of size N-1 such that it only
contains distinct integers in the range of 1 to N.
Find the missing element.
*/

/*
Sum of N integers from 1 to N is
(N * N+1)/2
 */

func (a ArrayOfIntegers) FindMissingNumberInGivenArray() int{
	sum := 0
	for _, val := range a.Arr{
		sum += val
	}

	// since we are missing one number from array the actual size is +1 the length of array
	actual_size := len(a.Arr)+1
	actual_sum := (actual_size * (actual_size+1))/2
	return actual_sum - sum
}
