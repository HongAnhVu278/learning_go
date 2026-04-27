package main

import "fmt"

// fib is a function that returns a function that return an int
func fib() func() int {
	// f(0) = 0
	// f(1) = 1
	// f(n) = f(n-1) + f(n-2)
	prev := 0
	curr := 1
	return func() int {
		ans := prev
		
		prev = curr
		curr = ans + curr

		return ans
			

	}

}

func main() {
	f := fib()

	for i:= 0; i < 10; i++ {
		fmt.Println(f())
	}
}