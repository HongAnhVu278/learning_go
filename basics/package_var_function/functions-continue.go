package main
import (
	"fmt"
)

// if two or more consecutive parameters share the same type => omit the the type from all but the last
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(12, 25))
}