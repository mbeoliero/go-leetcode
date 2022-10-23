/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	var res int
	left, right := 0, len(height)-1
	for left < right {
		temp := min(height[left], height[right]) * (right - left)
		if res < temp {
			res = temp
		}

		if min(height[left], height[right]) == height[right] {
			i := right - 1
			for ; i >= left && height[i] <= height[right]; i-- {
			}
			right = i
		} else {
			i := left + 1
			for ; i <= right && height[i] <= height[left]; i++ {
			}
			left = i
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

// @lc code=end

