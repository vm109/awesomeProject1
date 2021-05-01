package amazon

import (
	"fmt"
)

/*
In given number of numbers
example n = 3
1, 2, 3
the total number of set bits are
1 - 0001
2 - 0010
3 - 0011
so the total set bits are - 1+1+2 = 4
 */

func bit_representation_of_a_number(number int) int{
	var bit_representation []int
	total_ones := 0
	for number != 0 {
		//if(number == 2){
		//	bit_representation = append(bit_representation, 1)
		//	total_ones = total_ones + 1
		//	number = 0
		//}else {
			reminder := number % 2
			quotient := number / 2
			bit_representation = append(bit_representation, reminder)
			if(reminder == 1){
				total_ones = total_ones+ 1
			}

			number = quotient
		//}
	}

	fmt.Println(bit_representation)
	return total_ones

}
func Find_the_total_bits_for_given_number_of_numbers(numbers int){
	total_ones := 0
	i := 1
	for i <= numbers {
		total_ones = total_ones + bit_representation_of_a_number(i)
		i++
	}
	fmt.Println(total_ones)
}