func numTrees(n int) int {
	dp := make([]int, n+1)

	// Empty tree and single node tree
	dp[0] = 1
	dp[1] = 1

	for nodes := 2; nodes <= n; nodes++ {

		// Try every number as the root
		for root := 1; root <= nodes; root++ {

			left := root - 1
			right := nodes - root

			dp[nodes] += dp[left] * dp[right]
		}
	}

	return dp[n]
}