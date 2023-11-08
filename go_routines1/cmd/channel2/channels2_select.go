package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting a function to test select")

	// Create a Channel [ buffered or unbuffered ]
	ch := make(chan int)

	go func() {
		// sleep for 4 seconds and then send a message
		time.Sleep(4 * time.Second)
		ch <- 25
	}()

	select {
	case val := <-ch:
		fmt.Println(val)

	case <-time.After(5 * time.Second):
		panic("no message before 4 seconds")
	}
}
