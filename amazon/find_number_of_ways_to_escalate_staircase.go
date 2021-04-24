package amazon

/*
https://www.youtube.com/watch?v=5o-kdjv7FD0

If there are 4 steps
a person can go either 1 step or 2 steps
if A, B, C, D are steps

A to B to C to D
A to B to D
B to D
A to C to D


1 steps
a
1 step - 1 way

2 steps
a,b
1step - 1way
2nd step  - 2 ways
a,b
a
b
total ways 3 ways

3 steps
a - 1
b - 2
c - 2
d - 2
 */


func Travel_across_steps(total_steps int) int{

	if( total_steps > 2) {
		return Travel_across_steps(total_steps-1) + Travel_across_steps(total_steps-2)
	} else if ( total_steps == 1 ){
		return 1
	}else if ( total_steps == 2){
		return 2
	}
	return 0
}

func Num_Ways_Steps_Arr_Give( total_steps int, steps_allowed []int) int{
	if( total_steps > 2) {
		x := 0
		for _, value := range steps_allowed{
			x = x + Num_Ways_Steps_Arr_Give(total_steps-value, steps_allowed)
		}
		return x
	} else if ( total_steps == 1 ){
		return 1
	}else if ( total_steps == 2){
		return 2
	}
	return 0
}
