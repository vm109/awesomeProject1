package google

/*
Finding the first Duplicate Number in an array
 */
func FindFirstDuplicateNumber(arr []int)(duplicated int){
	duplicate_map := make(map[int]int)

	for _, value := range arr {
		if(duplicate_map[value] == 1 ){
			duplicated = duplicate_map[value]
			return
		}
		duplicate_map[value] = duplicate_map[value] + 1
	}
	return
}
