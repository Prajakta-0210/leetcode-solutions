func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for _, row := range matrix {
		for col := 0; col < cols; col++ {
			if row[col] == '1' {
				heights[col]++
			} else {
				heights[col] = 0
			}
		}

		area := largestRectangleArea(heights)

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		currentHeight := 0

		if i < len(heights) {
			currentHeight = heights[i]
		}

		for len(stack) > 0 &&
			heights[stack[len(stack)-1]] > currentHeight {

			heightIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[heightIndex]

			leftBoundary := -1
			if len(stack) > 0 {
				leftBoundary = stack[len(stack)-1]
			}

			width := i - leftBoundary - 1
			area := height * width

			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}