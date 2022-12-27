/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][3]int, n)
	// dp[i][0] 持有最大收益
	// dp[i][1] 不持有，但处于冷却期累计最大收益
	// dp[i][2] 不持有，处于非冷却期累计最大收益
	dp[0][0] = -prices[0]
	for i, price := range prices {
		if i == 0 {
			continue
		}
		// dp[i-1][0]: 继续持有，dp[i-1][2]-price: 买入
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-price)
		// 当天卖出股票(不持有)，变成冷却期
		dp[i][1] = dp[i-1][0] + price
		// 前一天卖出(dp[i-1][1])，前一天不操作(dp[i-1][2])
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

