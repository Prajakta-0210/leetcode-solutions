func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxWater := 0

	for left < right {
		width := right - left

		var area int
		if height[left] < height[right] {
			area = height[left] * width
			left++
		} else {
			area = height[right] * width
			right--
		}

		if area > maxWater {
			maxWater = area
		}
	}

	return maxWater
}