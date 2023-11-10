// context cancellation by timeout and request cancelling

package main

import (
	"context"
	"fmt"
	"time"
)

func executeFunctionTillTimeout(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("context is timedout")

	case <-time.NewTimer(4 * time.Second).C:
		fmt.Println("executed after 4 seconds")
	}
}

func main() {
	fmt.Println("starting signal timeout with context..!")
	ctx := context.Background()

	ctxT, _ := context.WithTimeout(ctx, 3*time.Second)

	executeFunctionTillTimeout(ctxT)
}
