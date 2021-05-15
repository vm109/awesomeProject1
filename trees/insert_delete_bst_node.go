package main

import "fmt"
/*

					  1			       	-----> 1
                           5			 ------> 2
						3		6		   --------> 3
					2	 	4				  --------> 4

if we want to delete node 11 then we need to make the node exchange with  - predecessor of the inorder traversal
or successor of the inorder traversal
 */

func attach_subtrees(subtree, attach_subtree *Node){
	if(subtree.value > attach_subtree.value){
		if(subtree.left != nil) {
			attach_subtrees(subtree.left, attach_subtree)
		}else{
			subtree.left = attach_subtree
		}
	}else if(subtree.value < attach_subtree.value){
		if(subtree.right != nil) {
			attach_subtrees(subtree.right, attach_subtree)
		}else{
			subtree.right = attach_subtree
		}
	}
}

func find_the_inorder_predeccessor(tree *Node, node_value int, previous *Node){
	if(tree.value < node_value && tree.right != nil){
		find_the_inorder_predeccessor(tree.right, node_value, tree)
	} else if(tree.value > node_value && tree.left != nil){
		find_the_inorder_predeccessor(tree.left, node_value, tree)
	}else if(tree.value == node_value && tree.left !=nil){
		temp := tree.right
		if(previous.value < tree.left.value ) {
			previous.right = tree.left
			attach_subtrees(previous.right, temp)
		}else if(previous.value > tree.left.value){
			previous.left = tree.left
		}
		fmt.Println(" ",tree.left.value)
	}else if( tree.left == nil){
		fmt.Println("there is no inorder predecessor")
	}
}

func insert_node_bst(tree *Node,  node_value int)  {
	if(node_value < tree.value){
		if(tree.left != nil) {
			insert_node_bst(tree.left, node_value)
		}else{
			tree.left = &Node{ value: node_value}
		}
	}else if(node_value > tree.value){
		if(tree.right != nil) {
			insert_node_bst(tree.right, node_value)
		}else{
			tree.right = &Node{ value: node_value}
		}
	}
}

func delete_node_bst(tree *Node, node_value int){
	if(tree.left != nil && tree.left.value == node_value){
		tree.left = nil
	}else if( tree.right != nil && tree.right.value == node_value){
		tree.right = nil
	}else if( tree.left != nil && tree.value > node_value){
		delete_node_bst(tree.left, node_value)
	}else if(tree.right != nil && tree.value < node_value){
		delete_node_bst(tree.right, node_value)
	}
}

func delete_node_with_one_child(tree *Node, value int){

}

func traverse_tree_inorder(tree *Node){
	if(tree.left != nil){
		traverse_tree_inorder(tree.left)
	}
	fmt.Print(tree.value)
	if(tree.right != nil){
		traverse_tree_inorder(tree.right)
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
