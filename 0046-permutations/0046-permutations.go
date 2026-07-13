func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {

		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])

			backtrack()

			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}