package main

import "fmt"

func main() {
	fmt.Println("main function is blocked without sender, i.e sender not available")

	// create a unbuffered channel of type int
	ch := make(chan int)

	// create a sender go routine
	// there will be a deadlock if we did not create a sender and still have receiver in main
	// go func() {
	// 	ch <- 26
	// }()

	// create a receiver
	// main function is deadlocked as there is a receiver without actually a sender to the channel
	val := <-ch

	fmt.Println(val)
}
