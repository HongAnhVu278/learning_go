// slice's like bookmarks on bookself: start, stop, shelf stop 
// start - stop: len
// start - shelf stop: cap
package main
import "fmt"

func main() {
	s := []int{1,3,5,7,9}
	printSlice(s)

	s = s[:0] // give it zero len
	printSlice(s)

	s = s[:4] // increase its len to 4
	printSlice(s)

	s = s[2:] // drop the first two values
	printSlice(s)


}

func printSlice(s []int) {
	// %d is strictly for integer 
	// %v is a "catch-all" for any value's default representation. 
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}