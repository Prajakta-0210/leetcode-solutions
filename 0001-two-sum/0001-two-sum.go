func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, num := range nums {
        complement := target - num

        if index, found := m[complement]; found {
            return []int{index, i}
        }

        m[num] = i
    }
    return []int{}
}