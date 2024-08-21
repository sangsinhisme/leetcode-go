package main

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
