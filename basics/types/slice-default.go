package main
import "fmt"

func main() {
	people := []struct{
		name string
		age int
	}{
		{"honganh", 23},
		{"ducanh", 28},
		{"ha", 25},
	} 
	
	sibling := people[:2]
	couple := people[1:]

	fmt.Println(sibling, couple)
}