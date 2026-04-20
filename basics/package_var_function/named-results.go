package main

import "fmt"

// named return
func split(sum int) (x, y int) {
	x = sum * 4
	y = sum - x
	return // "naked return" - returns x and y automatically => should only used in short function
}

func main() {
	fmt.Println(split(10))
}

/*
Unnamed returns:

func split(sum int) (int, int) {
	x := sum * 4
	y := sum - x

	return x, y

}
*/