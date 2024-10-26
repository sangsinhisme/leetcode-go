package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
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

func insertLevelOrder(arr []int, i int, n int) *TreeNode {
	var root *TreeNode
	if i < n {
		root = &TreeNode{Val: arr[i]}
		root.Left = insertLevelOrder(arr, 2*i+1, n)  // Left child
		root.Right = insertLevelOrder(arr, 2*i+2, n) // Right child
	}
	return root
}

func arrayToBinaryTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// Check if root is null
	if arr[0] == -1 {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	index := 1

	// Use a queue to insert nodes level by level
	for index < len(arr) {
		// Get the current node from the queue
		current := queue[0]
		queue = queue[1:]

		// Assign the left child
		if index < len(arr) && arr[index] != -1 {
			current.Left = &TreeNode{Val: arr[index]}
			queue = append(queue, current.Left)
		}
		index++

		// Assign the right child
		if index < len(arr) && arr[index] != -1 {
			current.Right = &TreeNode{Val: arr[index]}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	inorderTraversal(root.Right)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append(slices.Concat(postorderTraversal(root.Left), postorderTraversal(root.Right)), root.Val)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return append([]int{root.Val}, append(preorderTraversal(root.Left), preorderTraversal(root.Right)...)...)
}

func dfsTree(node *Node, post *[]int) {
	if node == nil {
		return
	}
	if len(node.Children) > 0 {
		for _, c := range node.Children {
			dfsTree(c, post)
		}
	}
	*post = append(*post, node.Val)
}

func postorder(root *Node) []int {
	var post []int
	dfsTree(root, &post)
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

/*
2807. Insert Greatest Common Divisors in Linked List
Given the head of a linked list head, in which each node contains an integer value.
Between every pair of adjacent nodes, insert a new node with a value equal to the greatest common divisor of them.
Return the linked list after insertion.
The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.
*/
func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	var helper func(head *ListNode) *ListNode
	helper = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}
		gcd := gcd(head.Val, head.Next.Val)
		head.Next = &ListNode{gcd, helper(head.Next)}
		return head
	}

	return helper(head)
}

/*
241. Different Ways to Add Parentheses
Given a string expression of numbers and operators, return all possible results from computing all the different
possible ways to group numbers and operators. You may return the answer in any order.
The test cases are generated such that the output values fit in a 32-bit integer and the number of different results
does not exceed 104.
*/
func diffWaysToCompute(expr string) []int {
	var result []int
	for i := 0; i < len(expr); i++ {
		char := expr[i]
		if char == '+' || char == '-' || char == '*' {
			left := diffWaysToCompute(expr[:i])
			right := diffWaysToCompute(expr[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch char {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}
	if len(result) == 0 {
		num, _ := strconv.Atoi(expr)
		result = append(result, num)
	}
	return result
}

/*
2583. Kth Largest Sum in a Binary Tree
https://leetcode.com/problems/kth-largest-sum-in-a-binary-tree/description/
*/
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var memo []int64
	var helper func(root *TreeNode, idx int)
	helper = func(root *TreeNode, idx int) {
		if root != nil {
			if idx >= len(memo) {
				memo = append(memo, int64(root.Val))
			} else {
				memo[idx] += int64(root.Val)
			}
			helper(root.Left, idx+1)
			helper(root.Right, idx+1)
		}
	}
	helper(root, 0)
	if k > len(memo) {
		return -1
	}
	sort.Slice(memo, func(i, j int) bool {
		return memo[i] > memo[j]
	})
	return memo[k-1]
}

/*
2641. Cousins in Binary Tree II
https://leetcode.com/problems/cousins-in-binary-tree-ii/description/
*/
func replaceValueInTree(root *TreeNode) *TreeNode {
	dp := make(map[int]int)
	var helper func(root *TreeNode, level int)
	helper = func(root *TreeNode, level int) {
		if root != nil {
			dp[level] += root.Val
			helper(root.Right, level+1)
			helper(root.Left, level+1)
		}
	}
	helper(root, 0)
	var replaceHelper func(root *TreeNode, level int, sibling int)
	replaceHelper = func(root *TreeNode, level int, sibling int) {
		if root != nil {
			root.Val = dp[level] - root.Val - sibling
			siblingLeft, siblingRight := 0, 0
			if root.Left != nil {
				siblingLeft = root.Left.Val
			}
			if root.Right != nil {
				siblingRight = root.Right.Val
			}
			replaceHelper(root.Right, level+1, siblingLeft)
			replaceHelper(root.Left, level+1, siblingRight)
		}
	}
	replaceHelper(root, 0, 0)
	return root
}

/*
951. Flip Equivalent Binary Trees
https://leetcode.com/problems/flip-equivalent-binary-trees/description/
*/
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		return (flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)) || (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right))
	} else {
		return false
	}
}

/*
2458. Height of Binary Tree After Subtree Removal Queries
https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/description/
*/
func treeQueries(root *TreeNode, queries []int) []int {
	n := len(queries)
	ans := make([]int, n)
	height := make(map[int]int)
	currHeight := 0
	var leftToRight func(node *TreeNode, level int)
	leftToRight = func(node *TreeNode, level int) {
		if node != nil {
			height[node.Val] = currHeight
			currHeight = max(currHeight, level)
			leftToRight(node.Left, level+1)
			leftToRight(node.Right, level+1)
		}
	}
	leftToRight(root, 0)
	currHeight = 0
	var rightToLeft func(node *TreeNode, level int)
	rightToLeft = func(node *TreeNode, level int) {
		if node != nil {
			height[node.Val] = max(height[node.Val], currHeight)
			currHeight = max(currHeight, level)
			rightToLeft(node.Right, level+1)
			rightToLeft(node.Left, level+1)
		}
	}
	rightToLeft(root, 0)
	for i, elem := range queries {
		ans[i] = height[elem]
	}
	return ans
}
