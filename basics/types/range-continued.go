// you can skip index, val by assigning _

package main

import "fmt"

func main() {
	nums := make([]int, 10)

	for i := range nums {
		nums[i] = 1 << uint(i)
	}

	for _, val := range nums {
		fmt.Printf("value: %d\n", val)
	}
}