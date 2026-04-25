//  append is built in function to grow a slice
// append(s []T, vs ...T) []T => ...T means any number of values of type T
/*
+) append ask the curr self "do i have room for more books on the right" (compare len to cap)
	-) if yes => place new book, bump len, return new bookmark
	-) if no => build a bigger shelf, copy everything oveer, plave new book => return bookmark for the new shelf

	=> old slice might still point to old shelf => if you forget s = > lose the appended data
*/

package main
import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s,1)
	printSlice(s)

	s = append(s,2,3,4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len:%d, cap: %d %v\n", len(s), cap(s), s)
}