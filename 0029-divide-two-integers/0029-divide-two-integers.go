func divide(dividend int, divisor int) int {

	// Overflow case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Determine sign
	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	// Convert to positive int64
	a := int64(dividend)
	if a < 0 {
		a = -a
	}

	b := int64(divisor)
	if b < 0 {
		b = -b
	}

	result := int64(0)

	for a >= b {

		temp := b
		multiple := int64(1)

		// Double divisor
		for a >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}

		a -= temp
		result += multiple
	}

	if sign < 0 {
		result = -result
	}

	return int(result)
}