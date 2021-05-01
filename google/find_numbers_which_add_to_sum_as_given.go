package google

/*
Find the matching pair of numbers in given collection of numbers
and the sum of the pair of numbers should be equal to a given sum

Source : https://www.youtube.com/watch?v=XKu_SEDAykw

Example :
1. [1,2,4,5] sum = 8 there is no pair
2. [1,3,4,4] sum = 8 pair of numbers [4,4]

 */

func Find_pair_of_numbers_from_collection(collection []int16, sum int16)[]int16{
	for i, value := range collection {
		temp := sum - value
		j := 0
		for j < len(collection){
			if( j!=i && collection[j] == temp){
				return []int16{collection[i], collection[j]}
			}
			j++
		}
	}
	return []int16{}
}

/*
This solution is for sorted array of numbers
 */
func Improved_Func_To_Find_Pair_Of_Numbers(collection []int16, sum int16) []int16{
	start_index := 0
	end_index := len(collection)-1
	for  int16(collection[end_index] )> sum && start_index != end_index {
		end_index = end_index -1
	}
	if( end_index == start_index){
		return []int16{}
	}
	for int16(collection[start_index]) != sum - int16(collection[end_index]) && start_index != end_index{
		start_index = start_index + 1
	}
	if( start_index == end_index){
		return []int16{}
	}else {
		return []int16{collection[start_index], collection[end_index]}
	}
}

/*
Improved solution with linear complexity
{2,5,18,6}  sum = 11
 */

func Finding_Pair_Of_Numbers_Linearly_With_Unsorted_Array(collection []int16, sum int16) []int16{
	complementary_map := map[int16]int{}
	for _, value := range collection{
		index_in_map := sum - value
		if(complementary_map[index_in_map] != 1){
			complementary_map[value] =1
		}else{
			return []int16{index_in_map, value}
		}
	}
	return []int16{}
}
