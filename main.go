package main

import "fmt"

func main() {
	root := []int{5, 4, 9, 1, 10, -1, 7}

	fmt.Println(preorderTraversal(replaceValueInTree(arrayToBinaryTree(root))))
}
