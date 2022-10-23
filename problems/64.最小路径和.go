/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	sum := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		sum += grid[i][0]
		dp[i][0] = sum
	}
	sum = 0
	for j := 0; j < n; j++ {
		sum += grid[0][j]
		dp[0][j] = sum
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

