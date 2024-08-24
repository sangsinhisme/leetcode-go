package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	x := 12389 - 12321
	y := 12389 - 12421
	fmt.Println(x, y)
	fmt.Println(nearestPalindromic("12389"))
}
