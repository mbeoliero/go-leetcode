/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
func candy(ratings []int) int {
	var ans int
	n := len(ratings)
	left, right := make([]int, n), make([]int, n)

	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 1
		}

		ans += max(left[i], right[i])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

