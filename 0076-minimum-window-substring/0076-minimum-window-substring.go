func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	required := len(t) 

	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		if need[char] > 0 {
			required--
		}

		need[char]--

		for required == 0 {
			windowLen := right - left + 1

			if windowLen < minLen {
				minLen = windowLen
				start = left
			}

			leftChar := s[left]
			need[leftChar]++

			if need[leftChar] > 0 {
				required++
			}

			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}