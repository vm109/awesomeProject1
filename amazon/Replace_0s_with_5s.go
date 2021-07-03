package amazon

import (
	"math"
)

/*
53426

53426/10 = 53420
 */
func find_zeros_in_integer(integer int32, arr []int32) []int32{
	if( int(integer)/10 !=0) {
		if(integer % 10 == 0) {
			arr = append(arr, 5)
		}else {
			arr = append(arr, integer%10)
		}
		arr = find_zeros_in_integer((integer) / 10, arr)
	}
	if(int(integer)/10 ==0 && integer> 0){
		arr = append(arr, (integer))
	}
	return arr
}
func ReplaceZerosWithFivesInAInteger(integer_number int32) int32{
	arr := make([]int32, 0)
	numbers_in_integer := find_zeros_in_integer(integer_number, arr)
	num := 0
	for i, val := range numbers_in_integer {
		num += int(float64(val)* math.Pow(10,float64(i)))
	}

	return int32(num)
}
