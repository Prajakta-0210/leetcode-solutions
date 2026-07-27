func isNumber(s string) bool {
	hasDigit := false
	hasDot := false
	hasExp := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case c >= '0' && c <= '9':
			hasDigit = true

		case c == '+' || c == '-':
			// Sign is only allowed at the beginning
			// or immediately after e/E
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}

		case c == '.':
			
			if hasDot || hasExp {
				return false
			}
			hasDot = true

		case c == 'e' || c == 'E':
		
			if hasExp || !hasDigit {
				return false
			}
			hasExp = true
			hasDigit = false 

		default:
			return false
		}
	}

	return hasDigit
}