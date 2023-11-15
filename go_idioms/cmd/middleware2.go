package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting a Server")
	http.ListenAndServe(":8080", loggerMiddleware(http.HandlerFunc(handler1)))
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next.ServeHTTP(w, r)
		fmt.Println("Request Processed in %v", time.Since(now))
	})
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello hi")
}
