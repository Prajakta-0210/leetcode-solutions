func reverse(x int) int {
	const (
		INT_MAX = 2147483647
		INT_MIN = -2147483648
	)

	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		// Check for positive overflow
		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > 7) {
			return 0
		}

		// Check for negative overflow
		if result < INT_MIN/10 || (result == INT_MIN/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}