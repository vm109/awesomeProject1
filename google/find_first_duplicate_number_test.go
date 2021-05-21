package google

import (
	"testing"
	assert "github.com/stretchr/testify/assert"
)

func TestFindingDuplicateNumber_When_Duplicate_Exists(t *testing.T){
		arr := []int{1, 2, 3, 4, 5, 6, 1}
		duplicate_number := FindFirstDuplicateNumber(arr)
		assert.Equal(t, duplicate_number, 1)
}

func TestFindDuplicateNumber_When_Duplicate_Doesnt_Exist(t *testing.T){
	arr := []int{1,4,5,7,8}
	duplicate_number := FindFirstDuplicateNumber(arr)
	assert.Equal(t, duplicate_number, 0)
}