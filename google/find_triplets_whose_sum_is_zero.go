package google

/*
{-4,-3,0,2,3,5,1,7}
1. we have split positives and negatives
2. [-4,-3] [0,2,3,5,1]
{0, -1, 2, -3, 1};
 */

func findTriplets(arr []int)map[int][]int{
	map1 := make(map[int]int)
	triplets := make(map[int][]int)
	triplet_arr_count := 0
	for index, value := range arr{
		for i := 0; i< len(arr); i++{
			if(map1[arr[i]] != 1){
				map1[arr[i]]  = 1
			}
			if( i != index) {
				sum := value + arr[i]
				if(map1[-(sum)] == 1 && (-(sum) != value && -(sum) != arr[i])){
					triplets[triplet_arr_count] = []int{value,-(sum), arr[i]}
					triplet_arr_count++
					map1[-(sum)] = 0
				}
			}
		}
	}
	return triplets
}

func FindTripletsInArrayWhoseSumIsZero(arr []int)map[int][]int{

	return findTriplets(arr)
}
