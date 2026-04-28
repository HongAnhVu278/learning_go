// method is a function with special "receiver" argument
/**
Go has no classes but methods give the same "object do things" feel
+) Java/Python: methods live inside the class definition
+) Go: it's defined outside the type, separately 
*/
package main
import (
	"fmt"
	"math"

)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3,4}

	fmt.Println(v.Abs())
}