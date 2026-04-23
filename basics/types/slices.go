// slice: dynamically sized, flexible view into the elements of the arr
// a[low :high] (low inclusice but high exclusive)

package main

import "fmt"

func main() {
	members := [4]string{"honganh","ha", "ducanh", "kor"}

	var human []string = members[0:3]
	fmt.Println(human)
}
