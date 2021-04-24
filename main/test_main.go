package main

import (
	amzn_questions "awesomeProject1/amazon"
	"fmt"
)
/*
this is a local file to run other functions in other packages
 */
//
//func main(){
//	first_non_repeating := amzn_questions.First_non_repeating_character_in_string("aaabbccddde")
//	fmt.Println(first_non_repeating)
//}


/*
Testing the print brackets for matrix amazon
 */

//func main(){
//	amzn_questions.Print_bracket_for_matrix_multiplication([]int{40, 20, 30, 10, 30}, 0)
//}


/*
Google question find the first recurring character in strung
string = helloo
string = helwitheworld
string = abc
 */
//
//func main()  {
//	str := "helloo"
//	repeating_character := google_questions.First_reccuring_charachter(str)
//	fmt.Println(repeating_character)
//	str = "helwitheworld"
//	repeating_character = google_questions.First_reccuring_charachter(str)
//	fmt.Println(repeating_character)
//	str = "abc"
//	repeating_character = google_questions.First_reccuring_charachter(str)
//	fmt.Println(repeating_character)
//}

/*
Function for finding total number of steps
 */

func main()  {
	total_ways := amzn_questions.Travel_across_steps(5)
	fmt.Println(total_ways)

	total_ways = amzn_questions.Num_Ways_Steps_Arr_Give(6, []int{1,3,5})
	fmt.Println(total_ways)
}