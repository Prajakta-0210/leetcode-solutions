func lengthOfLongestSubstring(s string) int {
    charMap := make(map[byte]int)

    left := 0
    maxLen := 0

    for right := 0; right < len(s); right++ {

        if index, found := charMap[s[right]]; found && index >= left {
            left = index + 1
        }

        charMap[s[right]] = right

        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }

    return maxLen
}