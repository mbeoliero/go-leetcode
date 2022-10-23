package problems

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var res int
	ori := x
	for {
		if x == 0 {
			break
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res == ori
}

// @lc code=end
