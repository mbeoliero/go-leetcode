/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] 删掉一个元素以后全为 1 的最长子数组
 */

// @lc code=start
func longestSubarray(nums []int) int {
	count := 0
	stack := make([]int, 0)
	for i, n := range nums {
		if n == 0 {
			if i == 0 {
				continue
			}
			if i > 0 && nums[i-1] == 1 {
				stack = append(stack, count)
				count = 0
			}
			if i > 0 && nums[i-1] == 0 {
				stack = append(stack, 0)
				count = 0
			}
		}
		if n == 1 {
			count++
		}
	}
	if count > 0 {
		stack = append(stack, count)
	}

	if count == len(nums) {
		return count - 1
	}
	if len(stack) == 1 {
		return stack[0]
	}

	ans := 0
	for i := 0; i < len(stack)-1; i++ {
		ans = max(ans, stack[i]+stack[i+1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

