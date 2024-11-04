package main

import (
	"container/heap"
	"math"
)

/*
2684. Maximum Number of Moves in a Grid
https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/description/
*/
func maxMoves(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	moves := [3][2]int{{-1, 1}, {0, 1}, {1, 1}}
	maxStep := 0

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var helper func(i, j int) int
	helper = func(i, j int) int {
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		maxMoves := 0
		for _, move := range moves {
			n1 := i + move[0]
			m1 := j + move[1]
			if n1 >= 0 && n1 < n && m1 >= 0 && m1 < m && grid[n1][m1] > grid[i][j] {
				maxMoves = max(maxMoves, helper(n1, m1)+1)
			}
		}

		memo[i][j] = maxMoves
		return maxMoves
	}

	for i := 0; i < n; i++ {
		maxStep = max(maxStep, helper(i, 0))
	}
	return maxStep
}

/*
778. Swim in Rising Water
https://leetcode.com/problems/swim-in-rising-water/description/
*/
func swimInWater(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	directions := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	minHeap := &MinHeapTime{{0, 0, grid[0][0]}}
	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		cur := heap.Pop(minHeap).(PointTime)
		i, j, time := cur.x, cur.y, cur.time
		for _, direct := range directions {
			nextI, nextJ := i+direct[0], j+direct[1]
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m {
				nextTime := max(grid[nextI][nextJ], time)
				if nextTime < dp[nextI][nextJ] {
					dp[nextI][nextJ] = nextTime
					heap.Push(minHeap, PointTime{nextI, nextJ, nextTime})
				}
			}
		}
	}
	return dp[n-1][m-1]
}
