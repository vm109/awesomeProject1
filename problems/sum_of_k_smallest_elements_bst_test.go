package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBSTRoot(t *testing.T) {
	r := CreateBSTRoot(20)
	root := &r
	root.AddRight(22)
	root.AddLeft(8)
	root = root.GoLeftNode()
	root.AddLeft(4)
	root.AddRight(12)
	root = root.GoRightNode()
	root.AddRight(14)
	root.AddLeft(10)

	fmt.Println(r)
}

func TestFindSmallestKNumbers(t *testing.T) {
	r := CreateBSTRoot(20)
	root := &r
	root.AddRight(22)
	root.AddLeft(8)
	root = root.GoLeftNode()
	root.AddLeft(4)
	root.AddRight(12)
	root = root.GoRightNode()
	root.AddRight(14)
	root.AddLeft(10)

	sum := FindSmallestKNumbers(&r, 3)
	assert.Equal(t, sum, 22)
}
