package main

import "fmt"

func main() {
	root := arrayToBinaryTree([]int{5, 8, 9, 2, 1, 3, 7, 4, 6})
	fmt.Println(treeQueries(root, []int{3, 2, 4, 8}))
}
