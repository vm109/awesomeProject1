package google

import "fmt"

/*
string = "helloo"
so the first recurring character is l
 */
func First_reccuring_charachter(given_string string) string {
	 var charachters []byte =  []byte(given_string)
	 character_map := map[byte]int{}
	 var temp string
	 for _, char := range charachters{
	 	fmt.Println(string(char), character_map[char])
		if(character_map[char] == 0){
			character_map[char] = 1
		}else if( character_map[char] ==1 ){
			temp = string(char)
			break
		}
	 }
	 return temp
}

