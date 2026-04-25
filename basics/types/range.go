// when ranging over a slice, two values are returned: index, copy of the element at that index

package main
import "fmt"

var nums = []int{1,3,5,7,9}

func main() {

	for i, v := range nums {
		fmt.Printf("index: %d, val: %d\n", i, v )
	}

}