package main

import (
	"bytes"
	"fmt"
)

/*
Maximum length of a concatenated string with unique characters
 */

/*
we need to try this with recursion
 */
func contains_chars(present_word, word_to_be_concatenated string) bool{
	if(len(present_word)>0 || len(word_to_be_concatenated)>0){
		return  bytes.ContainsAny([]byte(present_word), word_to_be_concatenated)
	}
	return  true
}

func find_big_unique_words(arr []string) int{
	preserve_string := ""
	for i,element := range arr{
		unique:= true
		for _, next_element := range arr[i+1:]{
			if(contains_chars(element, next_element)){
				if(len(element) > len(next_element)){
					preserve_string = preserve_string+element
				}
				unique = false
			}
		}
		if(unique){
			preserve_string = preserve_string + element
		}
	}
	return len(preserve_string)
}


func main(){
	arr := []string{"", "ab", "cd", "abcd", "cdab"}
	//length:= max_concatenated_string(arr)
	//fmt.Println(length)
	fmt.Println(find_big_unique_words(arr))
}
