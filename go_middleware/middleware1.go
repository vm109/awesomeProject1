package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("starting middleware1 function")
}

func main() {

	http.ListenAndServe(":8090", middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "- and in actual function")
	})))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "in middleware")
		next.ServeHTTP(w, r)
	})
}
