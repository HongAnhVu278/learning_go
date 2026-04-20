package main

import (
	"fmt"
	"math"
)

func pow(x,n,limit float64) float64 {
	if res := math.Pow(x,n); res < limit {
		return res
	} else {
		fmt.Printf("%g > %g\n", res, limit)
	}
	// can't use result here

	return limit
}

func main() {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}