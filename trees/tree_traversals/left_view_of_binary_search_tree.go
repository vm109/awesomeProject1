package tree_traversals

import (
	utils "awesomeProject1/my_utils"
	"awesomeProject1/trees/tree_structs"
	"fmt"
)
/*
Time complexity of this algoithm is O(n)
 */
type Traversal_Node struct {
	node *tree_structs.Node
	depth int
}

func bfsTraverseTree( queue utils.Queue, depth_previously_printed int){
	traverse_node := queue.Dequeue().(*Traversal_Node)
	if(traverse_node != nil) {
		if(traverse_node.depth != depth_previously_printed) {
			fmt.Println(traverse_node.node.Value)
			depth_previously_printed = traverse_node.depth
		}
		if (traverse_node.node.Left != nil) {

			queue.Enqueue(&Traversal_Node{traverse_node.node.Left, traverse_node.depth+1})
		}
		if(traverse_node.node.Right != nil) {
			queue.Enqueue(&Traversal_Node{traverse_node.node.Right, traverse_node.depth+1})
		}
	}
	if(len(queue.Queue) > 0){
		bfsTraverseTree(queue, depth_previously_printed)
	}
}

/*
	In Binary Search Tree, to print the left view of the tree
	traverse through nodes in BFS fashion and to the left node
	and print the first occurring node at that level
 */
func LeftViewOfBinarySearchTree(tree *tree_structs.Node){
	queue:= utils.Queue{}
	root_Node := &Traversal_Node{
		tree,
		0,
	}
	queue.Enqueue(root_Node)
	bfsTraverseTree(queue, -1 )
}
