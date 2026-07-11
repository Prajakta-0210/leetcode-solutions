func nextPermutation(nums []int) {
	n := len(nums)

	// Step 1: Find pivot
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 2: Find next greater element and swap
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 3: Reverse the suffix
	left := i + 1
	right := n - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}