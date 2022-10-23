/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// @lc code=start
type Pair struct {
	fir int // 先手
	sec int // 后手
}

func PredictTheWinner(nums []int) bool {
	return stoneGame(nums) >= 0
}

func stoneGame(nums []int) int {
	n := len(nums)
	dp := make([][]Pair, n)
	for i := range dp {
		dp[i] = make([]Pair, n)
		dp[i][i] = Pair{nums[i], 0} // 只有一个时，先手必胜
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := nums[i] + dp[i+1][j].sec  // 取了最左边的，对于右边堆是后手
			right := nums[j] + dp[i][j-1].sec // 取了最右边的，对于左边堆是后手

			// 取大的，后手选择根据取的值变化
			if left > right {
				dp[i][j].fir = left
				dp[i][j].sec = dp[i+1][j].fir
			} else {
				dp[i][j].fir = right
				dp[i][j].sec = dp[i][j-1].fir
			}
		}
	}

	res := dp[0][n-1]
	return res.fir - res.sec
}

// @lc code=end

