package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind_the_total_bits_for_given_number_of_numbers(t *testing.T) {
	binary_string := "10001"
	total_substrings := FindSubStringsWithStartAndEndingOnes(binary_string)
	assert.Equal(t, int32(1), total_substrings)
}


/*
so the picked ones are at indexes
0, 2, 4, 5
0 - 2
0 - 4
0 - 5

2 - 4
2 - 5

4 - 5
 */
func TestFind_the_total_bits_for_given_number_of_numbers_with_4_ones(t *testing.T) {
	binary_string := "1010101"
	total_substrings := FindSubStringsWithStartAndEndingOnes(binary_string)
	assert.Equal(t, int32(6), total_substrings)
}

func TestFind_the_total_bits_for_given_number_of_numbers_with_3_ones(t *testing.T) {
	binary_string := "00100101"
	total_substrings := FindSubStringsWithStartAndEndingOnes(binary_string)
	assert.Equal(t, int32(3), total_substrings)
}