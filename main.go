package main

import (
	_ "fmt"
)

func main() {
	// Values for the linked list
	values := []int{18}

	// Create the linked list
	head := createLinkedList(values)

	printLinkedList(insertGreatestCommonDivisors(head))
}
