// nil map = write clashes 
// reading from a nil map returns zero value but writing panics
// => make before write

package main
import "fmt"

type Info struct {
	Age int
	Job string
}

// key: string, value: Info
var m map[string]Info

func main() {
	m = make(map[string]Info)
	m["honganh"] = Info{
		23, "swe",
	}

	fmt.Println(m["honganh"])
}