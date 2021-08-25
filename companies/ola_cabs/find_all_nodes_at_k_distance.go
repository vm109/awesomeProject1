package ola_cabs

/*
Input: K = 0
		  1
		/   \
	   3     2
Output: 1
*/
func (t *Tree) FindAllNodesAtDistanceKFromRoot(height int) []int {
	q := &Queue{}
	q.Enqueue(t)
	return breathFirstSearch(q, height, 1)
}

func breathFirstSearch(q *Queue, height, number_of_nodes int) []int {
	// continue till we reach given height
	// or till we reach the end of tree
	present_height := int(FindHeightByNodes(number_of_nodes))
	arr := make([]int, 0)
	if !q.IsEmpty() {
		node := q.Dequeue().(*Tree)
		if present_height < height {
			number_of_nodes += 1
			if node != nil {
				// left side of the tree
				q.Enqueue(node.Left)
				// right side of the tree
				q.Enqueue(node.Right)
			} else {
				// left si e of the tree
				q.Enqueue((*Tree)(nil))
				// right side of the tree
				q.Enqueue((*Tree)(nil))
			}
		} else if present_height == height {
			if node != nil {
				arr = append(arr, node.Node)
			}
		}
		arr = append(arr,breathFirstSearch(q, height, number_of_nodes)...)
	}
	return arr
}
