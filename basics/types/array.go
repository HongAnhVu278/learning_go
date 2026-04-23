// [n]T => an array of n values of type T

package main
import "fmt"

func main() {
	var greeting[2] string
	greeting[0] = "hello"
	greeting[1] = "howdy"
	fmt.Println(greeting[0], greeting[1])
	fmt.Println(greeting)

	dogs := [3]string{"milo", "milu", "kor"}
	fmt.Println(dogs)

}