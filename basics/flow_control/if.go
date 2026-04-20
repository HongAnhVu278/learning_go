package main
import "fmt"

func greeting() {
	fmt.Println("hello")
}
func main() {
	for i:=0; i < 10; i++ {
		if i == 5 {
			greeting()
		} else {
			fmt.Println(i)
		}
	}
}