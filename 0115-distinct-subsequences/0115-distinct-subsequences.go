func numDistinct(s string, t string) int {
    m := len(s)
    n := len(t)

    // dp[j] = number of ways to form t[:j] 
    // using the characters processed from s
    dp := make([]int, n+1)

    // Empty t can always be formed in exactly 1 way
    dp[0] = 1

    for i := 1; i <= m; i++ {
        // Traverse backwards so that dp[j-1]
        // still represents the previous row.
        for j := n; j >= 1; j-- {
            if s[i-1] == t[j-1] {
                dp[j] = dp[j] + dp[j-1]
            }
        }
    }

    return dp[n]
}