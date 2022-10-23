/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var res int
	minP := prices[0]
	for _, p := range prices {
		res = max(res, p-minP)
		if minP > p {
			minP = p
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

