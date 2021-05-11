package search

func Linear_search(arr []int, value int) int{
	for i, val := range arr {
		if val == value {
			return i
		}
	}
	return -1
}
