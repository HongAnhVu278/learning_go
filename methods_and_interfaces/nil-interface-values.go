// a truly nil interface is one where both type and value are unset

package main
import "fmt"

type Speaker interface {
	Speak()
}

func main() {
	var s Speaker
	describe(s)

	// cause run-time err because there's no type in the interface tuple to indicate which concrete method to call
	s.Speak()

}

func describe(s Speaker) {
	fmt.Printf("value: %v | type: %T", s,s)

}