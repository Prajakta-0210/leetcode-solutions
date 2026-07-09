func myAtoi(s string) int {
	const (
		INT_MAX = 2147483647
		INT_MIN = -2147483648
	)

	i, n := 0, len(s)

	// Step 1: Skip leading spaces
	for i < n && s[i] == ' ' {
		i++
	}

	// Step 2: Check sign
	sign := 1
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	result := 0

	// Step 3: Read digits
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Step 4: Check overflow
		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > 7) {
			if sign == 1 {
				return INT_MAX
			}
			return INT_MIN
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}