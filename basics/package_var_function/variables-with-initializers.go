// var declaration can include initializer
// if initilizer is presented, type can be omitted
// car will tkae the type of initilizer

package main
import "fmt"

var num1, num2 int = 1, 2

func main() {
	var greeting, isClosed = "hello", true
	fmt.Println(num1, num2, greeting, isClosed)
}