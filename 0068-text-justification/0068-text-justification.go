
    
func fullJustify(words []string, maxWidth int) []string {
	var result []string
	n := len(words)
	i := 0

	for i < n {
		
		lineLen := len(words[i])
		j := i + 1

		for j < n && lineLen+1+len(words[j]) <= maxWidth {
			lineLen += 1 + len(words[j])
			j++
		}

		numWords := j - i
		lastLine := j == n

		wordLen := 0
		for k := i; k < j; k++ {
			wordLen += len(words[k])
		}

		var line strings.Builder

		if lastLine || numWords == 1 {
			for k := i; k < j; k++ {
				line.WriteString(words[k])
				if k != j-1 {
					line.WriteByte(' ')
				}
			}
		
			for line.Len() < maxWidth {
				line.WriteByte(' ')
			}
		} else {
			
			totalSpaces := maxWidth - wordLen
			gaps := numWords - 1
			evenSpaces := totalSpaces / gaps
			extraSpaces := totalSpaces % gaps

			for k := i; k < j; k++ {
				line.WriteString(words[k])

				if k != j-1 {
					spaces := evenSpaces
					if extraSpaces > 0 {
						spaces++
						extraSpaces--
					}
					line.WriteString(strings.Repeat(" ", spaces))
				}
			}
		}

		result = append(result, line.String())
		i = j
	}

	return result
}