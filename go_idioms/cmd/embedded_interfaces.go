package main

import (
	"fmt"
	"reflect"
)

type subinterface1 interface {
	Method1() bool
}

type subinterface2 interface {
	Method2() string
}

type number1 int
type number2 int

func (n number1) Method1() bool {
	fmt.Println("here in method1")
	fmt.Println(n)
	if reflect.ValueOf(n).Kind() == reflect.Int {
		return true
	}
	return false
}

func (n number2) Method2() string {
	fmt.Println("here in method2")
	fmt.Println(n)
	return reflect.ValueOf(n).String()
}

type superStruct struct {
	subinterface1
	subinterface2
}

func main() {
	fmt.Println("Starting Main Function for Embedded Interfaces")

	var one number1 = 1
	var two number2 = 2
	superStructVariable := &superStruct{
		subinterface1: one,
		subinterface2: two,
	}
	superStructVariable.Method1()
	superStructVariable.Method2()
}
