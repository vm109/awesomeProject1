package ola_cabs

import "fmt"

type ArrayOfIntegers struct {
	Arr []int
}

type Tree struct {
	Node  int
	Left  *Tree
	Right *Tree
}

type ArrayOfIntegerArrays struct {
	ArrOfArr [][]int
}

type Queue struct {
	Arr []interface{}
}

type QueueOperations interface {
	Enqueue(val interface{}) bool
	Dequeue() interface{}
	Seek() interface{}
	IsEmpty() bool
}

func (a ArrayOfIntegers) Contains(integer int) bool {
	contains := false
	for val := range a.Arr {
		if val == integer {
			contains = true
			break
		}
	}
	return contains
}

func (a ArrayOfIntegers) ContainsSameElementsThanGivenSet(set []int) bool {
	contains_same := true
	set_map := make(map[int]int, 0)
	for _, val := range set {
		set_map[val] = 1
	}

	for _, val := range a.Arr {
		if set_map[val] == 0 {
			contains_same = false
			break
		}
	}
	return contains_same
}

func (t *Tree) InOrderTraversalPrint() {
	if t != nil && t.Left != nil {
		t.Left.InOrderTraversalPrint()
	}
	fmt.Println(t.Node)
	if t != nil && t.Right != nil {
		t.Right.InOrderTraversalPrint()
	}
}

func (t *Tree) CompareTwoTreeToEqual(comparsion_tree *Tree) bool {
	left_tree_equal := false
	right_tree_equal := false
	both_equal := false
	if t != nil && t.Left != nil && comparsion_tree.Left != nil {
		left_tree_equal = t.Left.CompareTwoTreeToEqual(comparsion_tree.Left)
	} else if t.Left != nil && comparsion_tree.Left == nil {
		return false
	} else if t.Left == nil && comparsion_tree.Left == nil {
		left_tree_equal = true
	}

	both_equal = t.Node == comparsion_tree.Node

	if t != nil && t.Right != nil && comparsion_tree.Right != nil {
		right_tree_equal = t.Right.CompareTwoTreeToEqual(comparsion_tree.Right)
	} else if t.Right != nil && comparsion_tree.Right == nil {
		return false
	} else if t.Right == nil && comparsion_tree.Right == nil {
		right_tree_equal = true
	}

	return both_equal && left_tree_equal && right_tree_equal
}


func (q *Queue) IsEmpty() bool{
	return len(q.Arr) < 1
}

func (q *Queue) Enqueue(val interface{}) bool {
	if q!=nil {
		if q.Arr != nil {
			q.Arr = append(q.Arr, val)
		}else{
			q.Arr = make([]interface{},0)
			q.Arr = append(q.Arr, val)
		}
		return true
	}
	return false
}

func (q *Queue) Dequeue() interface{} {
	if(!q.IsEmpty()) {
		dequeueing_value := q.Arr[0]
		q.Arr = q.Arr[1:]
		return dequeueing_value
	}
	return nil
}

func (q *Queue) Seek() interface{}{
	if(!q.IsEmpty()){
		return q.Arr[0]
	}
	return nil
}