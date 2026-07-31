func largestRectangleArea(heights []int) int {
	stack := []int{} 
	maxArea := 0

	for i := 0; i <= len(heights); i++ {
		currentHeight := 0

		if i < len(heights) {
			currentHeight = heights[i]
		}

		for len(stack) > 0 && heights[stack[len(stack)-1]] > currentHeight {
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