func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	var dfs func(index, target int)

	dfs = func(index, target int) {

		if target == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		if target < 0 || index == len(candidates) {
			return
		}

		path = append(path, candidates[index])
		dfs(index, target-candidates[index]) 
		path = path[:len(path)-1]

		dfs(index+1, target)
	}

	dfs(0, target)

	return result
}