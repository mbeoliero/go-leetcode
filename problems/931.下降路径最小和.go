/*
 * @lc app=leetcode.cn id=931 lang=golang
 *
 * [931] 下降路径最小和
 */

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
			} else {
				dp[i][j] = matrix[i][j] + dp[i-1][j]
				if j == 0 {
					dp[i][j] = min(dp[i][j], matrix[i][j]+dp[i-1][j+1])
				} else if j == n-1 {
					dp[i][j] = min(dp[i][j], matrix[i][j]+dp[i-1][j-1])
				} else {
					dp[i][j] = min(dp[i][j], matrix[i][j]+dp[i-1][j+1])
					dp[i][j] = min(dp[i][j], matrix[i][j]+dp[i-1][j-1])
				}
			}

			if i == n-1 && j == 0 {
				res = dp[i][j]
			} else {
				res = min(res, dp[i][j])
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

