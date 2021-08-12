package ola_cabs

/*
Problem Statement:
Given an array of size N containing only 0s, 1s, and 2s; sort the array in ascending order.

Example:
Input:  {0 2 1 2 0}
Output: [0 0 1 2 2]
*/

func (a ArrayOfIntegers) SortArrayOf0S1SAnd2S() (sorted_arr []int){
	integer_map := make(map[int][]int, 0)
	// these are the expected integers set
	base_integers_set:= []int{0,1,2}
	if(!a.ContainsSameElementsThanGivenSet(base_integers_set)){

		return
	}
	for _, val := range a.Arr {
		if integer_map[val] == nil {
			arr := make([]int, 0)
			arr = append(arr, val)
			integer_map[val] = arr
		} else {
			integer_map[val] = append(integer_map[val], val)
		}
	}
	sorted_arr = make([]int, 0)

	for _, v := range base_integers_set{
		sorted_arr = append(sorted_arr, integer_map[v]...)
	}
	return
}
