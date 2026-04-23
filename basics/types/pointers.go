// house-address metaphor => each value has an address in the memory

package main

import "fmt"

func main() {
	i := 42
	
	
	p := &i // give me a piece of paper "p" with address if "i" on it
	fmt.Println(*p) // go to the house and look inside => print 42
	
	*p = 21 // go to the house and look inside + change the value inside to 21
	fmt.Println(i) // should print out 21

	j := 2701
	p = &j // p is now reassigned to j's address
	*p = *p/37 // change the value from 2701 to 73
	fmt.Println(j)
}