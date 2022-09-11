package main

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

func CreateBSTRoot(value int) (root Node) {
	root = Node{Value: value}
	return
}

func (n *Node) GoRightNode() (right *Node) {
	return n.RightNode
}

func (n *Node) GoLeftNode() (left *Node) {
	return n.LeftNode
}

func (n *Node) AddRight(value int) (root *Node) {
	n.RightNode = &Node{
		Value: value,
	}
	return n
}

func (n *Node) AddLeft(value int) (root *Node) {
	n.LeftNode = &Node{
		Value: value,
	}

	return n
}

func InOrderTraversal(n *Node, arr []int) []int {
	if n != nil {
		arr = InOrderTraversal(n.LeftNode, arr)
		arr = append(arr, n.Value)
		arr = InOrderTraversal(n.RightNode, arr)
	}
	return arr
}

func FindSmallestKNumbers(root *Node, K int) (smallNumbersSum int) {
	arr := make([]int, 0)
	arr = InOrderTraversal(root, arr)
	for i := 0; i < K; i++ {
		smallNumbersSum = smallNumbersSum + arr[i]
	}
	return
}
