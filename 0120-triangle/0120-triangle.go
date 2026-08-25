func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	// Copy the last row
	dp := make([]int, n)
	copy(dp, triangle[n-1])

	// Work from bottom to top
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}