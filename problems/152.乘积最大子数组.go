/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func maxProduct(nums []int) int {
	maxF, minF := make([]int, len(nums)), make([]int, len(nums))
	ans := nums[0]

	maxF[0], minF[0] = nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
		ans = max(ans, maxF[i])
	}
	return ans
}