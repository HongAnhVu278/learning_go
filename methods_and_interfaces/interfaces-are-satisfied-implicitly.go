// interfaces are implemented implicitly
// a type satisfies an interface just by having the right methods
// called structural typing (?) as opposed to nominal typing
// practical consequence: you can write an interface in package A and have a type in package B automatically satisfy it
// without even knowing A exists

package main
import "fmt"

// anyone can "speak"
type Speaker interface {
	Speak()
}

type Dog struct {
	Sound string
}

// we dont have to write "Dog implements Speaker"
// just by having a Speak() method, Dog automatically qualifies
func (d Dog) Speak()  {
	fmt.Println(d.Sound)
}

func main() {
	var s Speaker = Dog{"woof"}
	s.Speak()
}