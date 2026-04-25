/**
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. 

When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

package main

import (
	"golang.org/x/tour/pic"
	"fmt"
)

// return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integer
// the integers are interpreted as bluescale values, where the value 0 means full blue, and the value 255 means full white.
func Pic(dx, dy int) [][]uint8 {
	// create a slice of length dy
	outerSlice := make([][]uint8, dy)

	// create a loop to add slice of dx 8-bit
	for index, _ := range outerSlice {
		innerSlice := make([]uint8, dx)
		for i := 0; i < dx; i++ {
			val := (i+index)/2
			innerSlice[i] = uint8(val)
		}
		
		// outerSlice = append(outerSlice, innerSlice) 
		// when we create outerSlice in this make([][]uint8, dy) instead of make([][]uint8,0, dy)
		// it create sth like this: [[],[],[],[],[]] meaning [nil, nil, nil, nil, nil]
		// append does not replace nil slots, it add new slot at the end
		// get err: index out of range [0] with length 0 because of this

		outerSlice[index] = innerSlice
	}

	return outerSlice

}

func main() {
	result := Pic(5,5)
	fmt.Println(result)

	// Show display a picture defined by function f when excuted on Go playground
	pic.Show(Pic)

}


/**
NOTE:
imagine slice like an egg carton:
+) len: how many egg actually in the carton
+) capacity: how many slots the carton has in total

two way to make a slice:
+) make([]T, 5) => len=5, cap=5 => give me a carton with 5 eggs already in it 
=> you can gran egg #0, egg #1, etc
=> you can replace egg#2 : carton[2] = newEgg

+) make([]T, 0, 5) => len=0, cap=5 => give me an empty carton has space for 5 eggs
=> the carton exists, the slot exists but no egg
=> CANNOT carton[2] = newEgg 
=> you have to append

*/