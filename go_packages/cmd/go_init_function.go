package main

import (
	"fmt"
	testpack "packages/pkg"
)

func init() {
	fmt.Println("running init function")
}

func main() {
	fmt.Println("starting main function")
	testpack.TestPackFunction()
}
