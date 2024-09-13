package main

import (
	"fmt"
	_ "fmt"
)

func main() {

	fmt.Print(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}
