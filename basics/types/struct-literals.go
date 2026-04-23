// struct liter = "build a truct value inline"

package main

import "fmt"

type Person struct {
	Name string
	Age int
}

var (
	// cannot be p1 := Person{} because := cannot be used at package level
	p1 = Person{"honganh", 23} //positional
	p2 = Person{Name: "ducanh"} // named, partial
	p3 = Person{} //empty
	p = &Person{"kor", 2} //pointer literal: build and take address in one step => equivalent to p4 := Person{..} p = &p4
)

func main() {
	fmt.Println(p1, p2, p3, p)

}