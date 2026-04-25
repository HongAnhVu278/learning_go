package main
import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}

	board[0][0] = "X"
	board[0][2] = "X"
	fmt.Println(board[0])
	fmt.Println("------")
	
	board[1][2] = "X"
	board[1][0] = "O"
	fmt.Println(board[1])
	fmt.Println("------")

	board[2][2] = "O"
	fmt.Println(board[2])
	fmt.Println("------")
	

	for i :=0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}
}

