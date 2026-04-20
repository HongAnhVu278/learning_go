// := can be used inside function with implicit type
// oustide function, every statements begin with a key word(var, func, etc)

package main
import "fmt"

var num1, num2 int = 1, 2

func main() {
	var greeting string = "hello"
	isOpened := true
	isClosed := false

	fmt.Println(num1, num2, greeting, isOpened, isClosed)
}