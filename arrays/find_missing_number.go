package arrays

import "fmt"

/*
from 1 - 100 numbers
find a missing number.

Soltion:

const value is sum of 1 - 100 numbers
and we know the sum of all given numbers
then we will subtract sum of given numbers from sum of 100 numbers
 */


func Find_missing_numbers(given_array_of_numbers []int, n int){
	sum_of_100 := n*(n+1)/2
	sum := 0
	for _, n := range given_array_of_numbers{
		sum = sum + n
	}
	fmt.Println(sum_of_100 - sum)
}