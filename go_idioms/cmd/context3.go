package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done working")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	fmt.Println("cancelling a context with function")

	ctx := context.Background()

	ctxC, cancel := context.WithCancel(ctx)

	go worker(ctxC, 1)

	time.Sleep(7 * time.Second)

	fmt.Println("cancelling...")
	cancel()

	time.Sleep(2 * time.Second)
}
