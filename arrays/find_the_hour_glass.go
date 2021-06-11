package arrays

import "fmt"

func FindHourglassSum(arr [][]int32) int32 {
	// Write your code here
	var temp_sum int32 = -64
	for i:=0; i < len(arr)-2;i++ {
		for j:=0; j < len(arr[i])-2; j++ {
			sum := arr[i][j]+arr[i][j+1]+arr[i][j+2]+arr[i+1][j+1]+arr[i+2][j]+arr[i+2][j+1]+arr[i+2][j+2]
			fmt.Println(sum)
			fmt.Println(arr[i][j], arr[i][j+1], arr[i][j+2], arr[i+1][j+1], arr[i+2][j], arr[i+2][j+1], arr[i+2][j+2])
			if( sum > temp_sum){
				temp_sum = sum
			}
		}
	}
	return temp_sum
}