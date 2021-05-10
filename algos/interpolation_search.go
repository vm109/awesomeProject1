package algos
/*
Interpolation search:
In a sorted array, the time complexity is O(logN)
In sorted array jump search time complexity is O(sqrt(n))

Interpolation search is an improvement over binary search.
Where binary search will always look at the middle of the array
but Interpolation search will decide the direction of the search depending on
the value at middle.
 */
func Interpolation_search(arr []int, value, lo, hi int)int{

	if(lo <= hi && value >= arr[lo] && value <= arr[hi]){
		pos := lo + ((value - arr[lo])*(hi-lo)/(arr[hi]-arr[lo]))
		if(arr[pos] == value){
			return pos
		}else if( arr[pos]< value){
			return Interpolation_search(arr, value, pos+1, hi)
		}else if( arr[pos] > value){
			return Interpolation_search(arr, value, lo, pos-1)
		}
	}
	return -1
}
