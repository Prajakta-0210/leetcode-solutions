import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var result [][]int
	n := len(nums)

	for i := 0; i < n-3; i++ {

		// Skip duplicate first elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {

			// Skip duplicate second elements
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1

			for left < right {
				// Use int64 to avoid integer overflow
				sum := int64(nums[i]) + int64(nums[j]) +
					int64(nums[left]) + int64(nums[right])

				if sum == int64(target) {
					result = append(result, []int{
						nums[i],
						nums[j],
						nums[left],
						nums[right],
					})

					left++
					right--

					// Skip duplicate third elements
					for left < right && nums[left] == nums[left-1] {
						left++
					}

					// Skip duplicate fourth elements
					for left < right && nums[right] == nums[right+1] {
						right--
					}

				} else if sum < int64(target) {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}