package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	// Values for the linked list
	values := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}

	// Create the linked list
	head := createLinkedList(values)

	fmt.Println(spiralMatrix(3, 5, head))
}
