func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	top, bottom := 0, m-1
	left, right := 0, n-1

	var result []int

	for top <= bottom && left <= right {

		
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}