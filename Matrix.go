package main

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
