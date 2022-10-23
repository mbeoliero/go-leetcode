/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	zos := make([][]int, len(strs))
	for i, s := range strs {
		z, o := getZeroAndOne(s)
		zos[i] = []int{z, o}
	}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for k := 0; k < len(zos); k++ {
		for i := m; i >= 0; i-- {
			for j := n; j >= 0; j-- {
				// 倒序遍历
				if i >= zos[k][0] && j >= zos[k][1] {
					dp[i][j] = max(dp[i][j], dp[i-zos[k][0]][j-zos[k][1]]+1)
				} else {
					break
				}
			}
		}
	}
	return dp[m][n]
}

func getZeroAndOne(s string) (int, int) {
	zero, one := 0, 0
	for _, r := range s {
		if r == '0' {
			zero++
		}
		if r == '1' {
			one++
		}
	}
	return zero, one
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

