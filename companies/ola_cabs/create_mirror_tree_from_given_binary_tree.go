package ola_cabs

/*
Problem Statement:

Mirror of a Tree: Mirror of a Binary Tree T is another Binary Tree M(T) with left and right children
of all non-leaf nodes interchanged.
*/

func (t *Tree) ConvertTreeIntoItsMirrorViewTree() (*Tree){
	mirror_tree := &Tree{}
	treeMirrorByRecursion(t, mirror_tree)
	return mirror_tree
}

func treeMirrorByRecursion(original_tree, mirror_tree *Tree){
	// Node
	if(original_tree != nil){
		if(mirror_tree != nil) {
			mirror_tree.Node = original_tree.Node
		}else {
			mirror_tree = &Tree{
				Node:  original_tree.Node,
			}
		}
	}
	//Left Traversal
	if(original_tree != nil && original_tree.Left != nil) {
		mirror_tree.Right = &Tree{}
		treeMirrorByRecursion(original_tree.Left, mirror_tree.Right)
	}

	// Right Traversal
	if(original_tree != nil && original_tree.Right != nil) {
		mirror_tree.Left = &Tree{}
		treeMirrorByRecursion(original_tree.Right, mirror_tree.Left)
	}
}
