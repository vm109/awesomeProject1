package tree_misc

type TreeArr struct {
	Values []int
}
type Tree struct { 
	Value int
	Left *Tree
	Right *Tree
}

type TreeInterface interface {
	ValidateTreeIsBinary() bool
}

type TreeType string

const (
	LEFT_LINEAR_TREE TreeType = "left_linear_tree"
)

type TreeArrInterface interface {
	ConstructTreeFromArray(Type TreeType) Tree
}

func (ta TreeArr) ConstructTreeFromArray(Type TreeType) Tree {
	tree := Tree{}
	root := &tree
	switch Type {
	case "left_linear_tree":
		for i, value := range ta.Values{
			if(i == 0) {
				root.Value = value
			} else{
				root.Left = &Tree{ Value: value}
				root = root.Left
			}
		}
	}

	return tree
}


func (t Tree) ValidateTreeIsBinary() bool{

	return validateTreeIsBinary(&t)
}

func validateTreeIsBinary(t *Tree) bool{
	if((t.Left == nil && t.Right == nil)){
		return true
	}else if( t.Left != nil && t.Right !=nil) {
		validated := validateTreeIsBinary(t.Left)

		if(validated) {
			validated = validateTreeIsBinary(t.Right)
		}
		return validated
	}else{
		return false
	}
}