/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, i, j int) bool {
	n := 9
	if j == n {
		return backtrack(board, i+1, 0)
	}
	if i == n {
		return true
	}

	if board[i][j] != '.' {
		return backtrack(board, i, j+1)
	}

	for ch := '1'; ch <= '9'; ch++ {
		if !isValid(board, i, j, byte(ch)) {
			continue
		}

		board[i][j] = byte(ch)
		if backtrack(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, ch byte) bool {
	n := 9
	for k := 0; k < n; k++ {
		if board[i][k] == ch || board[k][j] == ch {
			return false
		}
		if board[(i/3)*3+k/3][(j/3)*3+k%3] == ch {
			return false
		}
	}
	return true
}

// @lc code=end

