package sort

/*
testing bubble sort
{5,4,1,10,9,3,13,20}
 */
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
