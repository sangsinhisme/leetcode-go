package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
1140. Stone Game II
Alice and Bob continue their games with piles of stones.  There are a number of piles arranged in a row, and each pile has a positive integer number of stones piles[i].  The objective of the game is to end with the most stones.
Alice and Bob take turns, with Alice starting first.  Initially, M = 1.
On each player's turn, that player can take all the stones in the first X remaining piles, where 1 <= X <= 2M.  Then, we set M = max(M, X).
The game continues until all the stones have been taken.
Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.
. . .
*/
func stoneGameII(piles []int) int {
	n := len(piles)
	sums := piles
	for i := n - 2; i > -1; i -= 1 {
		sums[i] += sums[i+1]
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n; i++ {
		dp[i][n] = sums[i]
	}
	for i := n; i > -1; i-- {
		for m := 1; m < n; m++ {
			var maxStones = 0
			maxStep := min(m*2, n-i)
			for x := 1; x <= maxStep; x++ {
				maxStones = max(maxStones, sums[i]-dp[i+x][max(m, x)])
			}
			dp[i][m] = maxStones
		}
	}
	return dp[0][1]
}

/*
*
947. Most Stones Removed with Same Row or Column

On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.
A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.
Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone, return the largest
possible number of stones that can be removed.
*/
func removeStones(stones [][]int) int {
	visited := make(map[[2]int]bool) // Use map for fast lookup
	islands := 0
	for _, elem := range stones {
		if _, ok := visited[[2]int{elem[0], elem[1]}]; !ok {
			removeLinkStones(stones, elem, visited)
			islands++
		}
	}
	return len(stones) - islands
}

func removeLinkStones(stones [][]int, stone []int, visited map[[2]int]bool) map[[2]int]bool {
	stack := Stack[[2]int]()
	stack.Push([2]int{stone[0], stone[1]})
	for stack.Length() > 0 {
		position := stack.Pop()
		visited[[2]int{position[0], position[1]}] = true
		for _, elem := range stones {
			if position[0] == elem[0] || position[1] == elem[1] {
				if !visited[[2]int{elem[0], elem[1]}] {
					stack.Push([2]int{elem[0], elem[1]})
				}
			}
		}
	}
	return visited
}

func maxScore(a []int, b []int) int {
	n := len(b)

	dp := make([][4]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			dp[i][j] = -999
		}
	}

	for i := 0; i < n; i++ {
		dp[i][0] = a[0] * b[i]
	}

	for j := 1; j < 4; j++ {
		bestPrev := math.MinInt32
		for i := 0; i < n; i++ {
			if i > 0 {
				bestPrev = max(bestPrev, dp[i-1][j-1])
			}
			if bestPrev != math.MinInt32 {
				dp[i][j] = max(dp[i][j], bestPrev+a[j]*b[i])
			}
		}
	}

	maxScore := math.MinInt32
	for i := 3; i < n; i++ {
		maxScore = max(maxScore, dp[i][3])
	}

	return maxScore
}
