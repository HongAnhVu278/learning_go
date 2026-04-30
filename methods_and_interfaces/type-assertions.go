/*
+) when you put a value into an interface, you put it into a box with two compartments(TYPE and VALUE)
=> to USE the value, you have to take it out of the box
=> type assertion is the syntax for taking it out: ex: i.(string)
*/
package main
import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)


}