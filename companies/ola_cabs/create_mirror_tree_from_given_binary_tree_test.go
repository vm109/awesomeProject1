package ola_cabs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_ConvertTreeIntoItsMirrorViewTree(t *testing.T) {
	original_tree := Tree{
		Node:  1,
	}
	original_tree.Left = &Tree{
		Node:  2,
	}
	original_tree.Right = &Tree{
		Node:  3,
	}

	mirror_tree := original_tree.ConvertTreeIntoItsMirrorViewTree()

	expected_mirror_tree := &Tree{
		Node:  1,
	}
	expected_mirror_tree.Left = &Tree{
		Node:  3,
	}
	expected_mirror_tree.Right = &Tree{
		Node:  2,
	}

	mirror_tree_validation := mirror_tree.CompareTwoTreeToEqual(expected_mirror_tree)
	assert.Equal(t, true, mirror_tree_validation)
}

func TestTree_ConvertTreeIntoItsMirrorViewTree_BiggerTree(t *testing.T) {
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

	mirror_tree := tree.ConvertTreeIntoItsMirrorViewTree()

	expected_mirror_tree := &Tree{
		Node: 1,
	}
	expected_mirror_tree.Left = &Tree{
		Node: 3,
	}
	expected_mirror_tree.Right = &Tree{
		Node: 2,
	}
	pointer = expected_mirror_tree.Right

	pointer.Right = &Tree{
		Node:  4,
	}

	pointer.Left = &Tree{
		Node: 5,
	}

	pointer = pointer.Right

	pointer.Left = &Tree{
		Node: 8,
	}

	pointer = expected_mirror_tree.Left

	pointer.Right = &Tree{
		Node: 6,
	}

	pointer.Left = &Tree{
		Node: 7,
	}
	are_two_trees_equal:= mirror_tree.CompareTwoTreeToEqual(expected_mirror_tree)
	assert.Equal(t, true, are_two_trees_equal)
}