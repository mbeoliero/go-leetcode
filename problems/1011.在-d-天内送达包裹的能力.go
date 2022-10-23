/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	left, right := 0, 1
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left < right {
		mid := left + (right-left)/2

		if function(weights, mid) == days {
			right = mid
		} else if function(weights, mid) < days {
			right = mid
		} else if function(weights, mid) > days {
			left = mid + 1
		}
	}
	return left
}

func function(weights []int, x int) int {
	var sum, days int
	for _, weight := range weights {
		if sum+weight > x {
			days++
			sum = 0
		}
		sum += weight
	}
	return days + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

