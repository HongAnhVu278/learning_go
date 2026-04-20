package main

import "fmt"

// var is used to declare multiple variables on one list
// name first, type last
var isOpened, isClosed bool

func main() {
	var number int
	fmt.Println(number, isOpened, isClosed)
}

/*
difference between var and :=
+) var: can be used inside a function with or without initial value
		can be used at package level (outside function)
		want to declare the type explicity
+) := : can only be used inside a function with initial value
		cannot ...... without initial value
		cannot be declared at package level
		cannot declare the type explicitly

*/