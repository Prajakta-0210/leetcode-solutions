func isScramble(s1 string, s2 string) bool {
	memo := make(map[string]bool)
	visited := make(map[string]bool)

	var dfs func(i1, i2, length int) bool

	dfs = func(i1, i2, length int) bool {
		key := fmt.Sprintf("%d,%d,%d", i1, i2, length)

		if visited[key] {
			return memo[key]
		}
		visited[key] = true

		if s1[i1:i1+length] == s2[i2:i2+length] {
			memo[key] = true
			return true
		}

		count := [26]int{}

		for i := 0; i < length; i++ {
			count[s1[i1+i]-'a']++
			count[s2[i2+i]-'a']--
		}

		for _, value := range count {
			if value != 0 {
				memo[key] = false
				return false
			}
		}

		for split := 1; split < length; split++ {
			if dfs(i1, i2, split) &&
				dfs(i1+split, i2+split, length-split) {
				memo[key] = true
				return true
			}

			if dfs(i1, i2+length-split, split) &&
				dfs(i1+split, i2, length-split) {
				memo[key] = true
				return true
			}
		}

		memo[key] = false
		return false
	}

	return dfs(0, 0, len(s1))
}