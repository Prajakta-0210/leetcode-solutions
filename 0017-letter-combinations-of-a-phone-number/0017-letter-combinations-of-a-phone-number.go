func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var backtrack func(index int, current string)

	backtrack = func(index int, current string) {
		// If we've processed all digits, store the combination.
		if index == len(digits) {
			result = append(result, current)
			return
		}

		letters := phone[digits[index]]

		for _, ch := range letters {
			backtrack(index+1, current+string(ch))
		}
	}

	backtrack(0, "")

	return result
}