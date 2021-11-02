package google

import tree "awesomeProject1/trees/tree_structs"

/*
Problem: Find the longest consecutive sequence in a binary tree.
The path refers to any sequence of nodes from some starting node to any node
Explanation:
                      1
                       \
                        3
                       /  \
                      2     4
                              \
                                5
in the above binary tree the longest sequence is [3,4,5]

The longest consecutive path need to be from parent to child and cannot be the reverse.
 */
func BinaryTreeLongestConsecutiveSequence(t *tree.Node, sequence_length int) int{
	ok, node := nodeHasBothChildren(t)
	if ok {
		sequence_length = 0
		sequence_length = BinaryTreeLongestConsecutiveSequence(t.Left, sequence_length)
		sequence_length = BinaryTreeLongestConsecutiveSequence(t.Right, sequence_length)
		return sequence_length
	}else if !ok && node!= nil {
		sequence_length += 1
		sequence_length = BinaryTreeLongestConsecutiveSequence(node, sequence_length)
		return sequence_length
	}else{
		if sequence_length > 1{
			sequence_length = sequence_length+ 1
		}
		return sequence_length
	}
}

func nodeHasBothChildren(n *tree.Node) (bool, *tree.Node) {
	if n.Right != nil && n.Left != nil {
		return true, nil
	}else if n.Right != nil && n.Left == nil{
		return false, n.Right
	}else if n.Left != nil && n.Right == nil{
		return false, n.Left
	}else {
		return false, nil
	}
}


/*
Algorithm after reading the code in geeks-for-geeks
 */
func BinaryTreeLongestConsecutiveNumberSequence(t *tree.Node, expected, sequence_length, max_length int) int{
	if t.Value!= expected {
		sequence_length = 1
		expected = t.Value + 1
	} else{
		expected += 1
		sequence_length += 1
	}

	if sequence_length > max_length {
		max_length = sequence_length
	}

	// go to left sub-tree
	if t.Left!= nil {
		max_length = BinaryTreeLongestConsecutiveNumberSequence(t.Left, expected, sequence_length, max_length)
	}

	// go to right sub-tree
	if t.Right != nil {
		max_length = BinaryTreeLongestConsecutiveNumberSequence(t.Right, expected, sequence_length, max_length)
	}
	return max_length
}