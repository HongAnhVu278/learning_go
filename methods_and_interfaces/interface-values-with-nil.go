package main
import "fmt"

type Printer interface {
	Print()
}

type Message struct {
	Text string
}

func (m *Message) Print() {
	if m == nil {
		fmt.Println("pointer is nil")
		return
	}

	fmt.Println(m.Text)
}

func main() {
	var p Printer

	// create a var called msg
	// its TYPE is *Message(pointer-to-Message) 
	// its VALUE is nil 
	var msg *Message

	// copies msg into interface p
	// the interace record same type-value thing
	p = msg
	describe(p)
	p.Print()

	p = &Message{"hello"}
	describe(p)
	p.Print()


}

func describe(p Printer) {
	fmt.Printf("value: %v | type: %T \n", p, p)
}

// interface ≠ nil even when the underlying value is nil, because the type word is set