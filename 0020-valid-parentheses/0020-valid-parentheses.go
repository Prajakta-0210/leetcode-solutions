func isValid(s string) bool {
	stack := []byte{}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]

		// Opening bracket
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			// Closing bracket but stack is empty
			if len(stack) == 0 {
				return false
			}

			// Check top of stack
			top := stack[len(stack)-1]

			if top != pairs[ch] {
				return false
			}

			// Pop
			stack = stack[:len(stack)-1]
		}
	}

	// Valid only if stack is empty
	return len(stack) == 0
}