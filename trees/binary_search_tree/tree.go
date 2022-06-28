package binary_search_tree

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewNode(value int) *Tree {
	return &Tree{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
func (t *Tree) Insert(node int) {
	if t != nil && node < t.Value {
		if t.Left != nil {
			t.Left.Insert(node)
		} else {
			t.Left = NewNode(node)
		}
	} else if t != nil && node > t.Value {
		if t.Right != nil {
			t.Right.Insert(node)
		} else {
			t.Right = NewNode(node)
		}
	}
}
