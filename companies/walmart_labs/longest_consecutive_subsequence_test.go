package walmart_labs

import (
	"fmt"
	"testing"
)

func TestLongestConsecutiveSubsequence(t *testing.T) {
	arr := []int{2,6,1,9,4,5,3}
	a := LongestConsecutiveSubsequence(arr)
	fmt.Println(a)
}

func TestLongestConsecutiveSubsequenceTwo(t *testing.T) {
	arr := []int{1,9,3,10,4,20,2}
	a := LongestConsecutiveSubsequence(arr)
	fmt.Println(a)
}