package amazon

import (
	"strconv"
)

func FindSubStringsWithStartAndEndingOnes(binary_string string) int32{
	ones_counter := 0
	binary_arr := []byte(binary_string)
	for _, value := range binary_arr {
		converted_int, _ := strconv.Atoi(string(value))
		if (converted_int) == 1 {
			ones_counter++
		}
	}
	total_substrings := 0
	if(ones_counter >= 2) {
		total_substrings = (ones_counter - 1) * (ones_counter) / 2
	}
	return int32(total_substrings)
}
