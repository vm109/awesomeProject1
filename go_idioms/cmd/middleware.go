package main

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func main() {
	fmt.Println("Testing Middleware in golang")

	limiter := NewRateLimiter(20, 4)
	server := http.Server{
		Addr:    ":8080",
		Handler: limiter.Middleware(http.HandlerFunc(handler)),
	}

	server.ListenAndServe()
}

type RateLimiter struct {
	limiter *rate.Limiter
}

func NewRateLimiter(rateLimit, burst int) (ratelimiter *RateLimiter) {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(rateLimit), int(burst)),
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !rl.limiter.Allow() {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello!")
}
