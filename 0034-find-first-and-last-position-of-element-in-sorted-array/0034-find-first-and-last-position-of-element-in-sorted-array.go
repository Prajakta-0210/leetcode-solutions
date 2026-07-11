func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)

	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			ans = mid
			right = mid - 1 // continue searching left
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1

	for left <= right {

		mid := left + (right-left)/2

		if nums[mid] == target {
			ans = mid
			left = mid + 1 // continue searching right
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}