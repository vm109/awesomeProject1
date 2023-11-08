package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("multiplexing reading out of channels")

	// create channel 1
	ch1 := make(chan int)

	// create channel 2
	ch2 := make(chan string)

	// create a sender for channel1
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 25
	}()

	// create a sender for channel2
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "hello world!"
	}()

	select {
	case ch1val := <-ch1:
		fmt.Println(ch1val)
	case ch2val := <-ch2:
		fmt.Println(ch2val)
	}
}
