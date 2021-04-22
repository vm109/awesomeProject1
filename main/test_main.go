package main

import (
	amzn_questions "awesomeProject1/amazon"
	"fmt"
)

/*
this is a local file to run other functions in other packages
 */

func main(){
	first_non_repeating := amzn_questions.First_non_repeating_character_in_string("aaabbccddde")
	fmt.Println(first_non_repeating)
}
