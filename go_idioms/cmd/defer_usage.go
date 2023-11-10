package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("worker number %d", i)
}

func deferUsageWithRoutines() {
	defer fmt.Println("done working on go routines")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}

	wg.Wait()
}

func deferWithPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic")
		}
	}()

	panic("boo panicked!")
}

func main() {
	fmt.Println("starting defer usage...!")

	deferUsageWithRoutines()

	deferWithPanic()
}
