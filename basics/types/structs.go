// struct is a collection of field => a custom type that groups the related data
// struct TYPE (the declaration) is like a class(the blueprint)
// struct VALUE(an instance: p := Vertex{1,2}) is like an obhect

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2})
}