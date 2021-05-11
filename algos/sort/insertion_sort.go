package sort

func Insertion_sort(arr []int) []int{
	for i, _ := range arr[:len(arr)-1] {
		if(arr[i+1]< arr[i]){
			swp := arr[i+1]
			j := i
			for j >=0 && !(arr[j] < swp) {
				arr[j+1] = arr[j]
				arr[j]=swp
				j--
			}
		}
	}
	return arr
}
