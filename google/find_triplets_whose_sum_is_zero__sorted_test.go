package google

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTripletsWithSorting2Triplet(t *testing.T) {
	arr := []int{9,2,5,8,-1,0,-8,3,6,-9,7}
	triplet_map := FindTripletsWithSorting(arr)
	assert.Equal(t, 2, len(triplet_map))
	assert.Equal(t, []int{0,-8,8}, triplet_map[0])
	assert.Equal(t, []int{0,-9,9}, triplet_map[1])
}

func TestFindTripletsWithSortingAnother2Triplet(t *testing.T){
	arr := []int{0, -1, 2, -3, 1}
	triplet_map := FindTripletsWithSorting(arr)
	assert.Equal(t, 2, len(triplet_map))
	assert.Equal(t, []int{0,-1,1}, triplet_map[0])
	assert.Equal(t, []int{1,-3,2}, triplet_map[1])
}

