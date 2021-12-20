package tree_misc

import "fmt"

func Find_Height_Of_Tree(tree Tree){
	// traverse through the tree.
	totalNodes := traverseThroughTheTree(&tree, 0)
	fmt.Println(totalNodes)
}

func traverseThroughTheTree(node *Tree, totalNodes int) int{
	if node != nil {
		totalNodes += 1
		totalNodes = traverseThroughTheTree(node.Left, totalNodes)
		totalNodes = traverseThroughTheTree(node.Right, totalNodes)
		return totalNodes
	}else {
		return totalNodes
	}
}
