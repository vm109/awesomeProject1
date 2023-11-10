package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Function calling in chain")
	ctx, cancel := context.WithTimeout(context.Background(), (4 * time.Second))
	defer cancel()
	caller1(ctx)
}

func caller1(ctx context.Context) {
	fmt.Println("in caller 1 function")
	caller2(ctx)
}

func caller2(ctx context.Context) {
	select {
	case <-time.After(4 * time.Second):
		fmt.Println("in caller2 function")
	case <-ctx.Done():
		fmt.Println("caller2: Timeout or cancellation detected")
	}
}
