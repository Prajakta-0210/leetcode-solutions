func findSubstring(s string, words []string) []int {

	if len(words) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalWords := len(words)
	totalLen := wordLen * totalWords

	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	var result []int

	for i := 0; i <= len(s)-totalLen; i++ {

		seen := make(map[string]int)

		j := 0

		for ; j < totalWords; j++ {

			start := i + j*wordLen

			word := s[start : start+wordLen]

			if _, ok := wordMap[word]; !ok {
				break
			}

			seen[word]++

			if seen[word] > wordMap[word] {
				break
			}
		}

		if j == totalWords {
			result = append(result, i)
		}
	}

	return result
}