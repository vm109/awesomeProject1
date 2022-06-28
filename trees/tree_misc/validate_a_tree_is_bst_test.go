package tree_misc

import (
	"testing"
)

func TestTreeArr_ConstructTreeFromArray(t *testing.T) {
	ta := TreeArr{
		Values: []int{5, 9, 1, 4},
	}
	tree := ta.ConstructTreeFromArray(LEFT_LINEAR_TREE)

	assert.Equal(t, tree.ValidateTreeIsBinary(), false)
}

func TestTree_ValidateTreeIsBinary_Testing_Non_Binary_Tree(t *testing.T) {
	tree := Tree{
		Value: 0,
	}
	tree.Left = &Tree{
		Value: 3,
	}
	tree.Right = &Tree{
		Value: 4,
	}

	tree.Left.Left = &Tree{
		Value: 5,
	}

	tree.Left.Right = &Tree{
		Value: 6,
	}

	//tree.Right.Left = &Tree{
	//	Value: 7,
	//}
	//
	tree.Right.Right = &Tree{
		Value: 8,
	}
	assert.Equal(t, tree.ValidateTreeIsBinary(), false)
}

func TestTree_ValidateTreeIsBinary_Testing_Binary_Tree(t *testing.T) {
	tree := Tree{
		Value: 0,
	}
	tree.Left = &Tree{
		Value: 3,
	}
	tree.Right = &Tree{
		Value: 4,
	}

	tree.Left.Left = &Tree{
		Value: 5,
	}

	tree.Left.Right = &Tree{
		Value: 6,
	}

	tree.Right.Left = &Tree{
		Value: 7,
	}

	tree.Right.Right = &Tree{
		Value: 8,
	}
	assert.Equal(t, tree.ValidateTreeIsBinary(), true)
}
