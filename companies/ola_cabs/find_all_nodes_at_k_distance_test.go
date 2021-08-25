package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_FindAllNodesAtDistanceKFromRoot(t *testing.T) {
	tree := &Tree{
		Node:  1,
	}
	tree.Left = &Tree{
		Node: 2,
	}
	tree.Right = &Tree{
		Node: 3,
	}
	pointer := tree.Left
	pointer.Left = &Tree{
		Node: 4,
	}
	pointer = pointer.Left
	pointer.Right = &Tree{
		Node: 8,
	}
	pointer = tree.Right
	pointer.Left = &Tree{
		Node: 5,
	}
	pointer.Right = &Tree{
		Node: 6,
	}
	nodes_at_level := tree.FindAllNodesAtDistanceKFromRoot(2)
	assert.Equal(t, []int{4, 5, 6}, nodes_at_level)

	nodes_at_level = tree.FindAllNodesAtDistanceKFromRoot(3)
	assert.Equal(t, []int{8}, nodes_at_level)
}

/*
	Input:
K = 3
        3
       /
      2
       \
        1
      /  \
     5    3
Output: 5 3

 */
func TestTree_FindAllNodesAtDistanceKFromRoot_Another_Tree(t *testing.T) {
	tree := &Tree{
		Node:  3,
	}
	tree.Left = &Tree{
		Node: 2,
	}

	pointer := tree.Left
	pointer.Right = &Tree{
		Node: 1,
	}

	pointer = pointer.Right
	pointer.Left = &Tree{
		Node: 5,
	}
	pointer.Right = &Tree{
		Node: 3,
	}

	nodes_at_level := tree.FindAllNodesAtDistanceKFromRoot(3)
	assert.Equal(t, []int{5, 3}, nodes_at_level)
}