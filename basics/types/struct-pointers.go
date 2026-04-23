/*
why pass a struct pointer instead of the struct itself:
+) change the original by "go to its address and write the value there"
+) cheaper: *Person carries ONE address (~8 bytes), no matter the struct size
   - Person (no *) copies every field into the function on each call
+) shortcut: mePointer.Name works even though mePointer is *Person
   - Go auto-dereferences the dot (no need to write (*mePointer).Name)
*/

package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	me := Person{"honganh", 23}

	mePointer := &me

	mePointer.Name = "anh"

	fmt.Println(me.Name)
}