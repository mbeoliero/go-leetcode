/*
 * @lc app=leetcode.cn id=775 lang=golang
 *
 * [775] 全局倒置与局部倒置
 */

// @lc code=start
func isIdealPermutation(nums []int) bool {
	n := len(nums)
	minSuf := nums[n-1]
	for i := n - 2; i > 0; i-- {
		if nums[i-1] > minSuf {
			return false
		}
		minSuf = min(minSuf, nums[i])
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

