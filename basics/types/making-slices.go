// create a slice using make
// a := make([]int, 5)  (len(a) = 5)
// b := make([]int, 0, 5) // len(b)=0, cap(b)=5

package main
import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	// under the hood, go allocates an arr of 5 int(all zero) but b's "window" in that arr has length 0  => []
	b := make([]int, 0, 5)
	printSlice("b", b)

	// give me a slice from index 0 to 2 of b's underlying array
	// => c points to the same underlying arr with len 2, cap 5
	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d) //len:3 cap: 3

	printSlice("b", b)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len:%d, cap:%d, %v \n",s, len(x), cap(x), x)
}