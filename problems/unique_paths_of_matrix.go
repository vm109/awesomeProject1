package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
m - rows
n - columns
size of the matrix M * N matrix
 */
func print_matrix_cells(m, n int){
	for i:=0; i<m; i++ {
		for j := 0 ; j< n; j++{
			fmt.Println(i,",",j)
		}
	}
}

//func reaching_down_in_matrix(m,n, max_row, max_column int){
//	if( m == max_row-1 && n == max_column-1){
//		fmt.Println(m,",",n)
//		fmt.Println("done reaching")
//	}else if(m < max_row || n < max_column ) {
//		fmt.Println(m,",",n)
//		if( m < max_row-1){
//			m++
//		}
//		if( n < max_column-1){
//			n++
//		}
//		reaching_down_in_matrix(m, n, max_row, max_column)
//	}
//}

func reach_next_step(x,y int, printer string){
	if(strings.Contains(printer, "->(0,0)")) {
		fmt.Println(printer )
	}
	new_x := x-1
	new_y := y-1

	if(new_x >= 0){

		reach_next_step(new_x, y, printer+"->("+strconv.Itoa(new_x)+","+strconv.Itoa(y)+")" )
	}
	if(new_y >= 0){
		reach_next_step(x, new_y, printer+"->("+strconv.Itoa(x)+","+strconv.Itoa(new_y)+")" )
	}
	if(new_x >= 0 && new_y >= 0){
		reach_next_step(new_x, new_y, printer+"->("+strconv.Itoa(new_x)+","+strconv.Itoa(new_y)+")" )
	}
}

