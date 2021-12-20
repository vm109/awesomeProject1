package amazon

/*
given array of integer nums and integer 'target', return 2 numbers from the array
that add upto the 'target' value

example:
arr := [1, 3, 4, 6, 2, 8]
target := 3

*/

/*
O(n^2) solution
*/
func ReturnTwoNumbersToAddValueToTarget(nums []int, target int) (pairnumbers []int) {
	pairnumbers = make([]int, 0)
	break_bool := false
	for i, val := range nums {
		if break_bool {
			break
		}
		if reminder := target - val; reminder > 0 {
			for _, val := range nums[i:] {
				if val == reminder {
					pairnumbers = append(pairnumbers, target-reminder)
					pairnumbers = append(pairnumbers, reminder)
					break_bool = true
					break
				}
			}
		}

	}
	return
}

/*
O(n) - hash solution
 */

func ReturnTwoNumbersFromTheArrayToAddToTarget_HashSolution(arr []int, target int) (pairnumbers []int) {
	complement_map := make(map[int]bool, 0)
	for _, value := range arr {
		target_complement := target - value
		if complement_map[target_complement] {
			return []int{target_complement, value}
		}else{
			complement_map[value] = true
		}
	}
	return
}
