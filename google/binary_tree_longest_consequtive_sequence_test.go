package google

import (
	"awesomeProject1/trees/tree_structs"
	"fmt"
	"testing"
)

func TestBinaryTreeLongestConsecutiveSequence(t *testing.T) {
	tree := tree_structs.Node{
		Value: 1,
	}
	tree.Right = &tree_structs.Node{
		Value: 3,
	}
	tree_child := tree.Right

	tree_child.Left = &tree_structs.Node{
		Value: 2,
	}
	tree_child.Right = &tree_structs.Node{
		Value: 4,
	}
	tree_child = tree_child.Right
	tree_child.Right = &tree_structs.Node{
		Value: 5,
	}

	sequence_length := BinaryTreeLongestConsecutiveSequence(&tree, 0)
	fmt.Println(sequence_length)
}

/*
Tree example:

   1
    \
     3
    /  \
   2     4
        /
       6
        \
		 7
		  \
           8


 */
func TestBinaryTreeLongestConsecutiveSequence_MoreTesting_One(t *testing.T) {
	tree := tree_structs.Node{
		Value: 1,
	}

	tree.Right = &tree_structs.Node{
		Value: 3,
	}

	tree_child := tree.Right

	tree_child.Left = &tree_structs.Node{
		Value: 2,
	}
	tree_child.Right = &tree_structs.Node{
		Value: 4,
	}

	tree_child = tree_child.Right

	tree_child.Left = &tree_structs.Node{
		Value: 6,
	}

	tree_child = tree_child.Left

	tree_child.Right = &tree_structs.Node{
		Value: 7,
	}

	tree_child = tree_child.Right

	tree_child.Right = &tree_structs.Node{
		Value: 8,
	}

	//sequence_length := BinaryTreeLongestConsecutiveSequence(&tree, 0)
	//fmt.Println(sequence_length)
	sequence_length := BinaryTreeLongestConsecutiveNumberSequence(&tree, tree.Value, 0, 1)
	fmt.Println(sequence_length)
}


/*
Tree example:

        1
       / \
      2   4
     /   / \
    3   5   6
           /
          7
*/
func TestBinaryTreeLongestConsecutiveSequence_MoreTesting_Two(t *testing.T) {
	tree := tree_structs.Node{
		Value: 1,
	}

	tree.Left = &tree_structs.Node{
		Value: 2,
	}
	tree.Right = &tree_structs.Node{
		Value: 4,
	}

	tree_child := tree.Left

	tree_child.Left = &tree_structs.Node{
		Value: 3,
	}

	tree_child = tree.Right

	tree_child.Left = &tree_structs.Node{
		Value: 5,
	}
	tree_child.Right = &tree_structs.Node{
		Value: 6,
	}

	tree_child = tree_child.Right

	tree_child.Left = &tree_structs.Node{
		Value: 7,
	}


	//sequence_length := BinaryTreeLongestConsecutiveSequence(&tree, 0)
	//fmt.Println(sequence_length)
	sequence_length := BinaryTreeLongestConsecutiveNumberSequence(&tree, tree.Value, 0, 1)
	fmt.Println(sequence_length)
}
