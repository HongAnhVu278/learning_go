package main

import "fmt"

type Info struct {
	Age int
	Job string
}

var m = map[string]Info {
	"honganh": Info {
		23, "swe",
	},
	"tien": Info {
		25, "data analyst",
	},
}

func main() {
	fmt.Println(m)
}