/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	res := 0
	maxFromLeft := make([]int, n)
	maxFromRight := make([]int, n)
	for i := 1; i < n; i++ {
		maxFromLeft[i] = max(maxFromLeft[i-1], height[i-1])
	}
	for j := n - 2; j >= 0; j-- {
		maxFromRight[j] = max(maxFromRight[j+1], height[j+1])
	}
	for i := 1; i < n-1; i++ {
		if height[i] < maxFromLeft[i] && height[i] < maxFromRight[i] {
			res += min(maxFromLeft[i], maxFromRight[i]) - height[i]
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

