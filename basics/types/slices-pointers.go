// slice does not store any data => it describes a section of underlying array
// change the element of slice modifies the corresponding element of the underlying arr

package main
import "fmt"

func main() {
	members := [5]string{
		"honganh",
		"ha",
		"ducanh",
		"kor",
		"milo",
	}
	fmt.Println(members)

	human := members[0:3]
	dogs := members[3:5]
	fmt.Println(human, dogs)

	dogs[0] = "milu"
	fmt.Println(human,dogs)
	fmt.Println(members)

}