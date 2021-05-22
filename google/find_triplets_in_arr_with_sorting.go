package google

func merge(left, right []int)[]int{
		i := 0;
		j := 0;
		sort_merge := []int{}
		for i < len(left) && j< len(right) && len(right) > 0 && len(left) >0{
			if(left[i] < right[j]){
				sort_merge = append(sort_merge, left[i])
				i++
			}else{
				sort_merge= append(sort_merge, right[j])
				j++
			}
		}
		for ( i < len(left)){
			sort_merge = append(sort_merge, left[i])
			i++
		}
		for( j < len(right)){
			sort_merge = append(sort_merge, right[j])
			j++
		}

	return sort_merge
}
func sortArray(arr []int)[]int{
	mid := len(arr)/2
	left := arr[0:mid]
	right := arr[mid:len(arr)]

	if(len(left) > 1) {
		left = sortArray(left)
	}
	if(len(right) > 1) {
		right = sortArray(right)
	}
	sorted_arr := merge( left, right)
	return sorted_arr
}

func findTripletsWithSorting(arr []int)map[int][]int{
	return_map := make(map[int][]int)
	triplet_arr_count := 0
	for index, value := range arr {
		i := 0
		j := len(arr)-1
		for i != j && i != index && j!=index{
			if(value+arr[i]+arr[j] == 0){
				return_map[triplet_arr_count] = []int{value, arr[i], arr[j]}
				triplet_arr_count++
			}
			i++
			j--
		}
	}
	return return_map
}

func FindTripletsWithSorting(arr []int)map[int][]int{
	sorted_arr := sortArray(arr)
	return findTripletsWithSorting(sorted_arr)
}