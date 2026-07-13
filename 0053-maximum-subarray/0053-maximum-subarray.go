func maxSubArray(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum < 0 {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}