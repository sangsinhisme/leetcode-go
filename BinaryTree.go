package main

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val      int
	Children []*Node
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(slices.Concat(postorderTraversal(root.Left), postorderTraversal(root.Right)), root.Val)
}

func dfs(node *Node, post *[]int) {
	if node == nil {
		return
	}
	if len(node.Children) > 0 {
		for _, c := range node.Children {
			dfs(c, post)
		}
	}
	*post = append(*post, node.Val)
}

func postorder(root *Node) []int {
	var post []int
	dfs(root, &post)
	return post
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var helper func(head *ListNode, root *TreeNode) bool
	helper = func(head *ListNode, root *TreeNode) bool {
		if root == nil {
			return false
		}
		if head == nil {
			return true
		}
		if root.Val == head.Val {
			return isSubPath(head.Next, root.Left) || isSubPath(head.Next, root.Right)
		}
		return false
	}

	if root == nil {
		return false
	}
	if helper(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

/*
2326. Spiral Matrix IV
You are given two integers m and n, which represent the dimensions of a matrix.
You are also given the head of a linked list of integers.
Generate an m x n matrix that contains the integers in the linked list presented in spiral order (clockwise),
starting from the top-left of the matrix. If there are remaining empty spaces, fill them with -1.
*/
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	visited := make(map[[2]int]bool)
	directions := map[int][2]int{
		0: {0, 1},  // Right
		1: {1, 0},  // Down
		2: {0, -1}, // Left
		3: {-1, 0}, // Up
	}
	direct := 0
	start := [2]int{0, 0}
	spiral := make([][]int, m)
	for i := range spiral {
		spiral[i] = make([]int, n)
		for j := range spiral[i] {
			spiral[i][j] = -1
		}
	}
	for head != nil {
		visited[start] = true
		spiral[start[0]][start[1]] = head.Val
		// next step
		vector := directions[direct]
		nextStart := [2]int{start[0] + vector[0], start[1] + vector[1]}
		if nextStart[0] < 0 || nextStart[1] < 0 || nextStart[0] > m-1 || nextStart[1] > n-1 || visited[nextStart] {
			direct = (direct + 1) % 4
			vector = directions[direct]
			nextStart = [2]int{start[0] + vector[0], start[1] + vector[1]}
		}
		start = nextStart
		head = head.Next
	}
	return spiral
}
