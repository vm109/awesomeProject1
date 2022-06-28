package tree_height_balance

import (
	"fmt"
	"testing"
	"trees/tree_structs"
)

func TestCheckIfTreeIsHeightBalanced(t *testing.T) {
	tree := &tree_structs.Node{
		Value: 2,
	}
	tree.Left = &tree_structs.Node{
		Value: 3,
		Left:  nil,
		Right: nil,
	}
	tree.Right = &tree_structs.Node{
		Value: 4,
		Left:  nil,
		Right: nil,
	}
	node := tree.Left
	node.Left = &tree_structs.Node{
		Value: 6,
	}
	node = node.Left
	node.Left = &tree_structs.Node{
		Value: 9,
	}

	balanced, height := CheckIfTreeIsHeightBalanced(tree, 0)
	fmt.Println(balanced, height)
}
