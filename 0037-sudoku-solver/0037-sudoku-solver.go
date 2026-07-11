func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			if board[row][col] == '.' {

				for ch := byte('1'); ch <= '9'; ch++ {

					if isValid(board, row, col, ch) {

						board[row][col] = ch

						if solve(board) {
							return true
						}

						// Backtrack
						board[row][col] = '.'
					}
				}

				return false
			}
		}
	}

	return true
}

func isValid(board [][]byte, row, col int, ch byte) bool {

	// Check row
	for i := 0; i < 9; i++ {
		if board[row][i] == ch {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == ch {
			return false
		}
	}

	// Check 3x3 box
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == ch {
				return false
			}
		}
	}

	return true
}