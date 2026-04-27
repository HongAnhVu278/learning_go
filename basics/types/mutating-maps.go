package main
import "fmt"

func main() {
	// create a map
	m := make(map[string]int)

	// insert an element in map
	m["answer"] = 42
	fmt.Println("the value:", m["answer"])
	// update an element in map
	m["answer"] = 47
	fmt.Println("the value after:", m["answer"])
	// delete the element
	delete(m, "answer")
	fmt.Println("after delete:", m["answer"])
	// test if a key is presented
	v, ok := m["answer"]
	fmt.Println("the value:", v, "presented?", ok)

}