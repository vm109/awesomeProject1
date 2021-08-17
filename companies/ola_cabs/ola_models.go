package ola_cabs

type ArrayOfIntegers struct {
	Arr []int
}

type Tree struct {
	Node int
	Left *Tree
	Right *Tree
}

type ArrayOfIntegerArrays struct {
	ArrOfArr [][]int
}

func (a ArrayOfIntegers) Contains(integer int) bool {
	contains := false
	for val := range a.Arr{
		if ( val == integer){
			contains = true
			break
		}
	}
	return contains
}

func (a ArrayOfIntegers) ContainsSameElementsThanGivenSet(set []int) bool{
	contains_same := true
	set_map := make(map[int]int, 0)
	for _, val := range set {
		set_map[val] = 1
	}

	for _, val := range a.Arr{
		if(set_map[val] == 0){
			contains_same = false
			break
		}
	}
	return contains_same
}