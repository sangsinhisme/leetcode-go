package main

import "slices"

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
