package google

import (
	"awesomeProject1/my_utils"
	"strings"
)

func GenerateBinaryStringFromGivenString(given_string string)[]interface{}{
	chars := []byte(given_string)
	queue := my_utils.Queue{
	}
	queue.Enqueue(given_string)
	for index, _ := range chars {
			if(len(queue.Queue)>0 && strings.Index(queue.Peek().(string),"?") == index) {
				a := queue.Dequeue()
				for a != nil && strings.Index(a.(string), "?") == index {
					new_string_bytes := []byte(a.(string))
					str0 := make([]byte, len(new_string_bytes))
					copy(str0, new_string_bytes)
					str0[index] = []byte("0")[0]
					queue.Enqueue(string(str0))
					str1 := make([]byte, len(new_string_bytes))
					copy(str1, new_string_bytes)
					str1[index] = []byte("1")[0]
					queue.Enqueue(string(str1))
					a = nil
					if(len(queue.Queue) > 0 && strings.Index(queue.Peek().(string),"?") == index){
						a = queue.Dequeue()
					}
				}
			}
	}
	return (queue.Queue)
}
