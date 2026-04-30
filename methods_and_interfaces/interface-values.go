package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name)
}

func main() {
	var s Speaker

	s = &Dog{Name: "Kor"}
	fmt.Printf("(%v, %T)\n", s, s) // %v: the value, %T: the type
	// print: (&{Kor}, *main.Dog)
	// &{Kor} : the value is a 'pointer to a struct containing its content: Kor'
	// *main.Dog : the type is 'pointer to Dog from package main'
	s.Speak()
}

// an interface var is a box with two compartments:
// type: *main.Dog
// value: &{Kor}