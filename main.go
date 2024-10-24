package main

import "fmt"

func main() {
	root1 := arrayToBinaryTree([]int{1, 2, 3, 4, 5, 6, -1, -1, -1, 7, 8})
	root2 := arrayToBinaryTree([]int{1, 3, 2, -1, 6, 4, 5, -1, -1, -1, -1, 8, 7})
	fmt.Println(flipEquiv(root1, root2))
}
