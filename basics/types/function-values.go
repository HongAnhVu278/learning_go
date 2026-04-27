// functions are valued and can be passed around like values
// func values may be used as function arguments and returns values
/** 
This is the foundation of Go's equivalent of "strategy
patterns." Instead of passing an interface with one method,
you often just pass the function itself. Common in sort
(sort.Slice(s, less)), HTTP handlers, error wrapping, etc
*/

package main
import (
	"fmt"
	"math"
)

//create a function  that takes in a func that return float64 as argument
func compute(fn func(float64,float64) float64) float64 {
	return fn(3,4)
} 

//in main
// create a function that takes in two float64 values that return math.Sqrt(x*x + y*y)

func main() {
	hypot := func(x,y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot)) // return hypot(3,4)
	fmt.Println(compute(math.Pow))
}

