package main

import (
	"container/heap"
	"sort"
)

/*
Q1. Find X-Sum of All K-Long Subarrays I
*/
func findXSumQ1(nums []int, k int, x int) []int {
	n := len(nums)
	freq := make(map[int]int)
	result := []int{}

	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}

	calculateXSum := func() int {
		type Pair struct {
			val  int
			freq int
		}
		pairs := []Pair{}
		for val, count := range freq {
			pairs = append(pairs, Pair{val, count})
		}

		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].freq == pairs[j].freq {
				return pairs[i].val > pairs[j].val
			}
			return pairs[i].freq > pairs[j].freq
		})

		// Sum the top x elements
		sum := 0
		for i := 0; i < x && i < len(pairs); i++ {
			sum += pairs[i].val * pairs[i].freq
		}
		return sum
	}

	result = append(result, calculateXSum())

	for i := k; i < n; i++ {
		freq[nums[i-k]]--
		if freq[nums[i-k]] == 0 {
			delete(freq, nums[i-k])
		}
		freq[nums[i]]++

		result = append(result, calculateXSum())
	}

	return result
}

/*
Q2. K-th Largest Perfect Subtree Size in Binary Tree
*/
func checkPerfectSubtree(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, 0, 0
	}
	leftPerfect, leftSize, leftDepth := checkPerfectSubtree(root.Left)
	rightPerfect, rightSize, rightDepth := checkPerfectSubtree(root.Right)

	if leftPerfect && rightPerfect && leftDepth == rightDepth {
		return true, leftSize + rightSize + 1, leftDepth + 1
	}
	return false, 0, 0
}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	var sizes []int

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		isPerfect, size, _ := checkPerfectSubtree(node)
		if isPerfect {
			sizes = append(sizes, size)
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	if len(sizes) < k {
		return -1
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[k-1]
}

/*
Q4. Find X-Sum of All K-Long Subarrays II
*/
type FreqElement struct {
	val  int
	freq int
}

type MinHeap []FreqElement

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].val < h[j].val
	}
	return h[i].freq < h[j].freq
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(FreqElement))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findXSum(nums []int, k int, x int) []int64 {
	freq := make(map[int]int)
	n := len(nums)
	result := []int64{}
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		freq[nums[i]]++
	}

	updateHeap := func() {
		h = &MinHeap{}
		heap.Init(h)
		for val, count := range freq {
			heap.Push(h, FreqElement{val, count})
		}
		for h.Len() > x {
			heap.Pop(h)
		}
	}

	calculateHeapSum := func() int64 {
		sum := int64(0)
		for _, elem := range *h {
			sum += int64(elem.val * elem.freq)
		}
		return sum
	}

	updateHeap()
	result = append(result, calculateHeapSum())

	for i := k; i < n; i++ {
		outgoing := nums[i-k]
		freq[outgoing]--
		if freq[outgoing] == 0 {
			delete(freq, outgoing)
		}

		incoming := nums[i]
		freq[incoming]++

		updateHeap()

		result = append(result, calculateHeapSum())
	}

	return result
}
