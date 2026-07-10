func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(curr string, open, close int)

	backtrack = func(curr string, open, close int) {
		// If the current string is complete
		if len(curr) == 2*n {
			result = append(result, curr)
			return
		}

		// Add '(' if possible
		if open < n {
			backtrack(curr+"(", open+1, close)
		}

		// Add ')' if it keeps the string valid
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return result
}