// a type switch is a switch where each case is a type (not a value)
// the special syntax i.(type) is only valid inside switch
// Type switches are common when handling JSON, RPC payloads, or anything that's "heterogeneous data" decoded into any. They give you
// both the type-test and the typed value in one move

package main
import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("twice %v is %v\n", v, 2*v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)	
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}