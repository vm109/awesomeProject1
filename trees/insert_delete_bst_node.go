package main

import (
	"awesomeProject1/trees/tree_structs"
	"fmt"
)
/*

					  1			       	-----> 1
                           5			 ------> 2
						3		6		   --------> 3
					2	 	4				  --------> 4

if we want to delete node 11 then we need to make the node exchange with  - predecessor of the inorder traversal
or successor of the inorder traversal
 */

func attach_subtrees(subtree, attach_subtree *tree_structs.Node){
	if(subtree.Value > attach_subtree.Value){
		if(subtree.Left != nil) {
			attach_subtrees(subtree.Left, attach_subtree)
		}else{
			subtree.Left = attach_subtree
		}
	}else if(subtree.Value < attach_subtree.Value){
		if(subtree.Right != nil) {
			attach_subtrees(subtree.Right, attach_subtree)
		}else{
			subtree.Right = attach_subtree
		}
	}
}

func find_the_inorder_predeccessor(tree *tree_structs.Node, node_value int, previous *tree_structs.Node){
	if(tree.Value < node_value && tree.Right != nil){
		find_the_inorder_predeccessor(tree.Right, node_value, tree)
	} else if(tree.Value > node_value && tree.Left != nil){
		find_the_inorder_predeccessor(tree.Left, node_value, tree)
	}else if(tree.Value == node_value && tree.Left !=nil){
		temp := tree.Right
		if(previous.Value < tree.Left.Value ) {
			previous.Right = tree.Left
			attach_subtrees(previous.Right, temp)
		}else if(previous.Value > tree.Left.Value){
			previous.Left = tree.Left
		}
		fmt.Println(" ",tree.Left.Value)
	}else if( tree.Left == nil){
		fmt.Println("there is no inorder predecessor")
	}
}

func insert_node_bst(tree *tree_structs.Node,  node_value int)  {
	if(node_value < tree.Value){
		if(tree.Left != nil) {
			insert_node_bst(tree.Left, node_value)
		}else{
			tree.Left = &tree_structs.Node{ Value: node_value}
		}
	}else if(node_value > tree.Value){
		if(tree.Right != nil) {
			insert_node_bst(tree.Right, node_value)
		}else{
			tree.Right = &tree_structs.Node{ Value: node_value}
		}
	}
}

func delete_node_bst(tree *tree_structs.Node, node_value int){
	if(tree.Left != nil && tree.Left.Value == node_value){
		tree.Left = nil
	}else if( tree.Right != nil && tree.Right.Value == node_value){
		tree.Right = nil
	}else if( tree.Left != nil && tree.Value > node_value){
		delete_node_bst(tree.Left, node_value)
	}else if(tree.Right != nil && tree.Value < node_value){
		delete_node_bst(tree.Right, node_value)
	}
}

func delete_node_with_one_child(tree *tree_structs.Node, value int){

}

func traverse_tree_inorder(tree *tree_structs.Node){
	if(tree.Left != nil){
		traverse_tree_inorder(tree.Left)
	}
	fmt.Print(tree.Value)
	if(tree.Right != nil){
		traverse_tree_inorder(tree.Right)
	}
}

//func main(){
//	root := Node{
//		value: 1,
//	}
//	root.right = &Node{
//		value: 5,
//	}
//	traverse_tree_inorder(&root)
//	fmt.Println("--------------------")
//	insert_node_bst(&root, 3)
//	insert_node_bst(&root, 6)
//	insert_node_bst(&root, 4)
//	insert_node_bst(&root, 2)
//	traverse_tree_inorder(&root)
//	find_the_inorder_predeccessor(&root, 5, nil)
//	fmt.Println("---------------------")
//	traverse_tree_inorder(&root)
//}
