package tree_traversals

import (
	"awesomeProject1/trees/tree_structs"
	"fmt"
)

func inorder_traverse(root *tree_structs.Node, value int, arr []int)[]int{
		if (root.Left != nil) {
			arr = inorder_traverse(root.Left, value, arr)
		}
		if (root.Value == value) {
			if (root.Left != nil) {
				arr = append(arr, root.Left.Value)
			}
			if (root.Right != nil) {
				arr = append(arr, root.Right.Value)
			}
		}
		if (root.Right != nil) {
			arr = inorder_traverse(root.Right, value, arr)
		}
	return arr
}

func Inorder_predecessor_successor(root *tree_structs.Node, value int)[]int{
	arr := inorder_traverse(root, value, []int{})
	return arr
}

func Inorder_predecssor_successor_tester(root *tree_structs.Node, arr []int){
	for _, value := range arr {
		predecessor_successor := Inorder_predecessor_successor(root, value)
		fmt.Println("predecessor successor values for ", value, "start ....")
		for _, pred_suc_val := range predecessor_successor {
			fmt.Println( pred_suc_val)
		}
		fmt.Println("predecessor successor values for ", value, "done ....")
	}
}
