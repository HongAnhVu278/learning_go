// every go file must start with package declaration
// this tells the Go compiler that this package build into an executable program not a library
package main

// fmt is the standard lib (short for format) - which handle formatted I/O
import "fmt"

// in a package main, a function main is the program's entry point
func main() {
	// call Println function from fmt package
	fmt.Println("hello world")
}

// when you run go run hello.go, go compiles code into a binary(machine-executable code) and the OS run it
