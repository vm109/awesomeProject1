package tree_height_balance

import (
	"math"
	structs "trees/tree_structs"
)

func CheckIfTreeIsHeightBalanced(tree *structs.Node, nodeLevel int) (balanced bool, height int) {
	leftTreeBalanced := false
	rightTreeBalanced := false
	leftSubtreeHeight := nodeLevel
	rightSubTreeHeight := nodeLevel
	if tree.Left != nil {
		leftTreeBalanced, leftSubtreeHeight = CheckIfTreeIsHeightBalanced(tree.Left, nodeLevel+1)
	} else {
		leftTreeBalanced = true
	}

	if tree.Right != nil {
		rightTreeBalanced, rightSubTreeHeight = CheckIfTreeIsHeightBalanced(tree.Right, nodeLevel+1)
	} else {
		rightTreeBalanced = true
	}

	if leftTreeBalanced && rightTreeBalanced && int(math.Abs(float64(leftSubtreeHeight)-float64(rightSubTreeHeight))) <= 1 {
		if leftSubtreeHeight == rightSubTreeHeight && rightSubTreeHeight == 0 {
			return true, nodeLevel
		}
		if leftSubtreeHeight > rightSubTreeHeight {
			return true, leftSubtreeHeight
		} else {
			return true, rightSubTreeHeight
		}
	} else {
		if leftSubtreeHeight == rightSubTreeHeight && rightSubTreeHeight == 0 {
			return false, nodeLevel
		}
		if leftSubtreeHeight > rightSubTreeHeight {
			return false, leftSubtreeHeight
		} else {
			return false, rightSubTreeHeight
		}
	}
}
