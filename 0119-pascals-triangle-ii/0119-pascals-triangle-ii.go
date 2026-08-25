func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i <= rowIndex; i++ {
		// Traverse backwards so previous values aren't overwritten
		for j := i; j >= 1; j-- {
			row[j] = row[j] + row[j-1]
		}
	}

	return row
}