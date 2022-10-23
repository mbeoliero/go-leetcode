/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
	end, far := 0, 0
	jumps := 0

	for i := 0; i < n-1; i++ {
		far = max(far, nums[i]+i)
		if end == i {
			jumps++
			end = far
		}
	}

	return jumps
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

func jump(nums []int) int {
	return dp(nums, 0)
}

func dp(nums []int, p int) int {
	if p >= len(nums)-1 {
		return 0
	}

	res := math.MaxInt16
	for i := 1; i <= nums[p]; i++ {
		res = min(res, dp(nums, p+i)+1)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}