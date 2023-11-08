package main

import "fmt"

func main() {
	fmt.Println("starting a dead-lock situation")

	// creating a unbuffered channel of type int
	ch := make(chan int)

	// spinning a go routine to receive data out of channel
	// now without a reciver there is a deadlock
	// go func() {
	// 	val := <-ch
	// 	fmt.Println(val)
	// }()

	// sender is blocked until a receiver is available
	ch <- 26
}
