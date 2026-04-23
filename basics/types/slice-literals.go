//slice literal is arr literal without lenght

package main
import "fmt"

func main() {
	numbers := []int{1,2,3}
	fmt.Println(numbers)

	boo := []bool{true, false, false}
	fmt.Println(boo)

	people := []struct {
		name string
		age int
	}{
		{"honganh", 23},
		{"ducanh", 28},
		{"ha", 25},
	}
	fmt.Println(people)
}