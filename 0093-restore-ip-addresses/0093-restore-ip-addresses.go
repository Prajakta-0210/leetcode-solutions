func restoreIpAddresses(s string) []string {
	result := []string{}
	parts := []string{}

	var backtrack func(start int)

	backtrack = func(start int) {
		if len(parts) == 4 {
			if start == len(s) {
				result = append(result, strings.Join(parts, "."))
			}
			return
		}

		remainingChars := len(s) - start
		remainingParts := 4 - len(parts)

		if remainingChars < remainingParts ||
			remainingChars > remainingParts*3 {
			return
		}

		for length := 1; length <= 3 && start+length <= len(s); length++ {
			part := s[start : start+length]

			if len(part) > 1 && part[0] == '0' {
				break
			}

			value := 0
			for i := 0; i < len(part); i++ {
				value = value*10 + int(part[i]-'0')
			}

			if value > 255 {
				break
			}

			parts = append(parts, part)
			backtrack(start + length)
			parts = parts[:len(parts)-1]
		}
	}

	backtrack(0)

	return result
}