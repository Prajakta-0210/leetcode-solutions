func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var path []int

	var dfs func(start, target int)

	dfs = func(start, target int) {

		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i < len(candidates); i++ {

			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > target {
				break
			}

			path = append(path, candidates[i])

			dfs(i+1, target-candidates[i])

			path = path[:len(path)-1]
		}
	}

	dfs(0, target)

	return result
}