package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Example:
	Tree:

          1
       /     \
     2        3
   /   \    /   \
  4     5   6    7
   \
     8

In The above tree left view of the tree is 1, 2, 4, 8
*/

func TestTree_LeftViewOfTree_BaseCase(t *testing.T) {
	tree := Tree{
		Node:  1,
	}
	tree.Left = &Tree{
		Node:  2,
	}
	tree.Right = &Tree{
		Node:  3,
	}
	pointer := tree.Left

	pointer.Left = &Tree{
		Node:  4,
	}

	pointer.Right = &Tree{
		Node: 5,
	}

	pointer = pointer.Left

	pointer.Right = &Tree{
		Node: 8,
	}

	pointer = tree.Right

	pointer.Left = &Tree{
		Node: 6,
	}

	pointer.Right = &Tree{
		Node: 7,
	}

	arr := tree.LeftViewOfTree([]int{})
	assert.Equal(t, []int{1,2,4,8}, arr)
}

func TestTree_LeftViewOfTree_Small_Binary_Tree(t *testing.T) {
	tree := Tree {
		Node: 1,
	}
	tree.Left = &Tree{
		Node: 3,
	}
	tree.Right = &Tree{
		Node: 2,
	}

	arr := tree.LeftViewOfTree([]int{})
	assert.Equal(t,[]int{1,3}, arr )
}