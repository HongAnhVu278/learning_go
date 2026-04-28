// method's receiver can be a value (v Vertex) or pointer (v *Vertex)
// a value receiver hands the method a copy of the struct
// a pointer receiver hands it the keys to the actual value
// => pointer receivers are the only way to mutate the receiver
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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
} 

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}