package main

import (
	"fmt"
	_ "fmt"
	"math"
)

func minSteps(n int) int {
	var factors = 0
	for n%2 == 0 {
		n = n / 2
		factors += 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			n = n / i
			factors += i
		}
	}

	if n > 2 {
		factors += n
	}
	return factors
}

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

func main() {

	var n = []int{1, 2, 3, 4, 5, 100}
	fmt.Println(stoneGameII(n))
}
