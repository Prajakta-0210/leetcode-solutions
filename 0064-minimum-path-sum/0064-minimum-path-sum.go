func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	// First row
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// First column
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// Fill remaining cells
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[m-1][n-1]
}