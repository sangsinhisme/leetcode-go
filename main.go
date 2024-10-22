package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 8, 9, 2, 1, 3, 7, 4, 6}
	n := len(arr)
	root := insertLevelOrder(arr, 0, n)
	fmt.Println(kthLargestLevelSum(root, 2))
}
