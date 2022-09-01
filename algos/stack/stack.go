package stack

type Stack struct {
	Arr []int
}

type StackI interface {
	Get() int
	Add(val int)
	Peek() int
	Length() int
}

/*
GET function returns the first value from the stack and removes it from the array
*/

func (s *Stack) Get() (val int) {
	if s.Length() <= 0 {
		return -1
	} else {
		len := len(s.Arr)
		val = s.Arr[len-1]
		s.Arr = s.Arr[:len-1]
	}
	return
}

/*
Add function will allow to add a value into the stack.
*/

func (s *Stack) Add(val int) {

	if s.Arr != nil {
		s.Arr = append(s.Arr, val)
	} else {
		s.Arr = make([]int, 0)
		s.Arr = append(s.Arr, val)
	}
}

/*
Peek will allow to view the latest value which is added into the stack
*/

func (s *Stack) Peek() (val int) {
	if len(s.Arr) <= 0 {
		return -1
	} else {
		len := len(s.Arr)
		val = s.Arr[len-1]
	}
	return
}

/*
Length will measure the length of the underlying array. This gives the size of values in the stack
*/

func (s *Stack) Length() (length int) {

	if s.Arr != nil {
		length = len(s.Arr)
	} else {
		return -1
	}
	return
}
