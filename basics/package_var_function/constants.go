/*
const: 
+) values computed at compile time (unlike var - computed at runtime)
+) allowed type: bool, number, string, char 
*/

package main

import "fmt"

const Pi = 3.14

func main() {
	const Name = "honganh"

	fmt.Println("hello", Name)
	fmt.Println("happy", Pi, "day")
}