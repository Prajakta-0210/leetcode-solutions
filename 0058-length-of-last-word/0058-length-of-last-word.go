func lengthOfLastWord(s string) int {
    n := len(s)
    count := 0

    // Skip trailing spaces
    i := n - 1
    for i >= 0 && s[i] == ' ' {
        i--
    }

    // Count characters of the last word
    for i >= 0 && s[i] != ' ' {
        count++
        i--
    }

    return count
}