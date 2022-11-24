/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] 找到最高海拔
 */

// @lc code=start
func largestAltitude(gain []int) int {
	res := 0
	sum := 0
	for _, g := range gain {
		sum += g
		if sum > res {
			res = sum
		}
	}

	return res
}

// @lc code=end

