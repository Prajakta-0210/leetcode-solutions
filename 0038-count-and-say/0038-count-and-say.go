func countAndSay(n int) string {
	result := "1"

	for i := 2; i <= n; i++ {
		var sb strings.Builder

		count := 1

		for j := 1; j < len(result); j++ {

			if result[j] == result[j-1] {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(result[j-1])
				count = 1
			}
		}

		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(result[len(result)-1])

		result = sb.String()
	}

	return result
}