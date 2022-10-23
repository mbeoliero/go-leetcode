/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if farthest <= i {
			return false
		}
	}
	return farthest >= len(nums)-1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

