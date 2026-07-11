func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m := len(num1)
	n := len(num2)

	result := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			mul := int(num1[i]-'0') * int(num2[j]-'0')

			p1 := i + j
			p2 := i + j + 1

			sum := mul + result[p2]

			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}

	var sb strings.Builder

	start := 0

	for start < len(result) && result[start] == 0 {
		start++
	}

	for ; start < len(result); start++ {
		sb.WriteByte(byte(result[start] + '0'))
	}

	return sb.String()
}