package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}

	root.Left = node2
	node2.Left = node3
	node3.Left = node4
	fmt.Println(postorderTraversal(root))
}
