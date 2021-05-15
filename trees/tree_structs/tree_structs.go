package tree_structs

type Tree struct {
	Height int
	Nodes int
}

type Node struct {
	Value int
	Left *Node
	Right *Node
}
