package sort

/*
testing bubble sort
{5,4,1,10,9,3,13,20}
 */

type Array struct{
	Arr []interface{}
}
func Bubble_sort(arr []int) []int{
	i := 0
	for i < len(arr) {
		j:=0
		for j+1 < len(arr)-i{
			if(arr[j+1] < arr[j]){
				swap_number := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = swap_number
			}
			j++
		}
		i++
	}
	return arr
}


/*
A Bubble sort has time complexity of O(n2)
Because in Bubble sort we are iterating over the array 2 times
 */
func (a Array) BubbleSortInteger() []interface{}{
	i := 0
	for i < len(a.Arr)-1 {
		j := i+1
		for j < len(a.Arr) {
			if a.Arr[i].(int) > a.Arr[j].(int){
				t := a.Arr[i]
				a.Arr[i] = a.Arr[j]
				a.Arr[j] = t
			}
			j++
		}
		i++
	}
	return a.Arr
}

/*
Insertion Sort - Insertion sort is similar to sorting playing cards.
we have a part of the array which is sorted and another part of the array unsorted.

start with index 1 and compare it with the predecessor.

Time complexity of insertion sort is still O(n2) since we have to walk through the array twice 
 */

func (a Array) InsertionSort() []interface{} {
	i := 1
	for range a.Arr[i:] {
		for j, _ := range a.Arr[:i] {
			if a.Arr[j].(int) > a.Arr[i].(int){
				swp := a.Arr[j]
				a.Arr[j] = a.Arr[i]
				a.Arr[i] = swp
			}
		}
		i++
	}
	return a.Arr
}


/*
Merge Sort - merge sort uses divide and conquer method
 */

func (a Array) MergeSort() []interface{}{
	return  split(a.Arr)
}

func split(arr []interface{}) []interface{}{
	var leftArray []interface{}
	var rightArray []interface{}
	if len(arr) > 2 {
		initial := 0
		final := len(arr) - 1
		mid := (initial + final) / 2
		leftArray = arr[initial:mid]
		rightArray = arr[mid:]
	} else{
		leftArray = []interface{}{arr[0]}
		rightArray = []interface{}{arr[1]}
	}
	if len(leftArray) > 1 {
		leftArray = split(leftArray)
	}
	if len(rightArray) > 1 {
		rightArray = split(rightArray)
	}
	return merge(leftArray, rightArray)
}

func merge(leftArray, rightArray []interface{}) []interface{}{
	a := make([]interface{}, 0)
	minLength := 0
	if len(leftArray) > len(rightArray){
		minLength = len(rightArray)
	}else {
		minLength = len(leftArray)
	}
	i := 0
	j := 0
	for i < minLength && j < minLength {
		if leftArray[j].(int) < rightArray[i].(int) {
			a = append(a, leftArray[j])
			j++
		} else {
			a = append(a, rightArray[i])
			i++
		}
	}

	for j < len(leftArray) {
		a = append(a, leftArray[j])
		j++
	}

	for i < len(rightArray) {
		a = append(a, rightArray[i])
		i++
	}
	return a
}