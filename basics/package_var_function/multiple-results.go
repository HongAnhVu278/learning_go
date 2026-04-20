package main

import "fmt"

// Go allows multiple values returned in one function
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// := do two things at once: declare new var and infer its type from right hand side
	// shorthand for:
	// var a string
	// var b string
	// a, b = swap("hello", "world")
	x, y := swap("hello", "world")
	fmt.Println(x, y)
}