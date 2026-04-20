package main

import "fmt"

func main() {
	// the function call is not excuted until the surrounding functions returned
	defer fmt.Println("world")

	fmt.Println("hello")
}