// an interface type is defined as a set of method signatures
// a var declared with an interface type can store any value
// as long as that value's type has the methods the interface requires

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}
func (d Dog) Speak() string {
	return "Woof!"
}

type Robot struct{}
func(r *Robot) Speak() string {
	return "Beep beep"
}

func main() {
	var s Speaker

	dog := Dog{}
	robot := Robot{}

	s = dog
	fmt.Println(s.Speak())

	s = &robot // s hold the address of robot
	// s = robot would not compile
	fmt.Println(s.Speak())
}