/**
Implement WordCount. 
It should return a map of the counts of each “word” in the string s. 
The wc.Test function runs a test suite against the provided function and prints success or failure.
*/

package main
import (
	"golang.org/x/tour/wc"
	"strings"
	//"fmt"
)

func WordCount(s string) map[string]int {
	// initialize the result map 
	res := make(map[string]int)

	// split the string into an arr of words in s
	words := strings.Fields(s)

	// count the frequency of each word in s and add in map
	for _, word := range words {
		res[word]++
	}

	return res

}

func main() {

	wc.Test(WordCount)
}