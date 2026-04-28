// you can declare a method on non-struct type too
// => "domain types" in Go. For ex: instead of passing around raw ints for user IDs,
// you might define type UserID int64 and give it methods

package main
import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}