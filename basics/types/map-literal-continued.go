package main

import "fmt"

type Info struct {
	Age int
	Name string
}

var m = map[string]Info {
	"honganh": {23, "swe"},
	"tien": {25, "data analyst"},
}

func main() {
	fmt.Println(m)
}