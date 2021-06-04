package main


func jumpingOnClouds(c []int32) int32 {
	canbewoninstep := 0
	// Write your code here
	i := 0
	if(1 < len(c) && len(c) < 101 && c[0] == 0 && c[len(c)-1] == 0 ) {
		for i < len(c)-1 {
			if( c[i] > 1 || c[i] < 0) {
				canbewoninstep = 0
				break
			}
			if( c[i+1] == 0 ) {
				canbewoninstep = canbewoninstep + 1
				i = i + 2
			}else if ( c[i] == 0){
				canbewoninstep = canbewoninstep + 1
				i = i + 1
			}else{
				canbewoninstep = 0
				break
			}
		}
	}
	return int32(canbewoninstep)
}

func main(){

}
