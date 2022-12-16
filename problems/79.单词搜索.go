/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])

	op := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}

		if board[i][j] == word[idx] {
			ch := board[i][j]
			board[i][j] = '.'
			for _, o := range op {
				if dfs(i-o[0], j-o[1], idx+1) {
					return true
				}
			}
			board[i][j] = ch
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

// @lc code=end

