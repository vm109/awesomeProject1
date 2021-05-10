package main

import (
	"awesomeProject1/algos"
	"fmt"
)

//func main(){
//	arr := []int{2,7,8,9,19,10}
//
//	index:= algos.Linear_search(arr, 19)
//	if(index <0){
//		fmt.Println("number not found")
//	}else{
//		fmt.Println("number found at position ", index)
//	}
//}

//func main(){
//	arr:= []int{1,2,3,5,8,10,12,15,19}
//	index := algos.Binary_search(arr, 19, 0, 8)
//	if(index < 0){
//		fmt.Println("number not found")
//	}else {
//		fmt.Println(index)
//	}
//}

//func main(){
//	arr := []int{1,2,3,4,5,6,7,8,9,10}
//	index := algos.Jump_search(arr, 1)
//	fmt.Println(index)
//}

func  main()  {
	arr:= []int{3,9,81,6561,43046721}
	index := algos.Interpolation_search(arr, 43046721, 0, len(arr)-1)
	fmt.Println(index)
}