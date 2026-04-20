package main
import (
	"fmt"
	"math"
)
	

func pow(x, n, limit float64) float64 {
	if res := math.Pow(x, n); res < limit {
		return res
	}
	return limit
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
		pow(2,1,1),
	)

}