package tree_misc

import "testing"

func TestFind_Height_Of_Tree(t *testing.T) {
	tree := Tree{
		Value: 1,
	}

	tree.Left = &Tree{
		Value: 2,
	}
	tree.Right = &Tree{
		Value: 3,
	}
	node := tree.Left
	node.Left = &Tree{
		Value: 4,
	}

	Find_Height_Of_Tree(tree)
}
