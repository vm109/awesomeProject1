package amazon

/*
non repeating character -
"aaabccc" - b is the first non repeating character
"abcd" - a is the first non repeating character
 */
func First_non_repeating_character_in_string(given_string string) string{
	var non_repeating_char byte
	char_count := 0
	for i,char := range []byte(given_string){
		if( i== 0){
			non_repeating_char = char
			char_count++
		}else if( non_repeating_char != char){
			if(char_count == 1){
				return string(non_repeating_char)
			}
			char_count = 1
			non_repeating_char = char
		}else{
			char_count++
		}
	}
	return "0"
}


