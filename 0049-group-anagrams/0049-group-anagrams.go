func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		mp[key] = append(mp[key], str)
	}

	result := [][]string{}
	for _, group := range mp {
		result = append(result, group)
	}

	return result
}