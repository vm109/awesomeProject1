package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_CompareTwoTreeToEqual_False(t *testing.T) {
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
		Node: 4,
	}
	pointer.Right = &Tree{
		Node: 5,
	}

	compare_tree := Tree{
		Node:  1,
	}
	compare_tree.Left = &Tree{
		Node:  2,
	}
	compare_tree.Right = &Tree{
		Node:  3,
	}
	pointer = compare_tree.Left
	pointer.Left = &Tree{
		Node: 5,
	}
	pointer.Right = &Tree{
		Node: 4,
	}
	both_trees_equal := tree.CompareTwoTreeToEqual(&compare_tree)
	assert.Equal(t, false, both_trees_equal)
}

func TestTree_CompareTwoTreeToEqual_True(t *testing.T) {
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
		Node: 4,
	}
	pointer.Right = &Tree{
		Node: 5,
	}

	compare_tree := Tree{
		Node:  1,
	}
	compare_tree.Left = &Tree{
		Node:  2,
	}
	compare_tree.Right = &Tree{
		Node:  3,
	}
	pointer = compare_tree.Left
	pointer.Left = &Tree{
		Node: 4,
	}
	pointer.Right = &Tree{
		Node: 5,
	}
	both_trees_equal := tree.CompareTwoTreeToEqual(&compare_tree)
	assert.Equal(t, true, both_trees_equal)
}
