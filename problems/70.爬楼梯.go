/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a, b := 1, 1
	for i := 2; i <= n; i++ {
		// c := a + b
		// a = c
		// b = a
		a, b = a+b, a
	}
	return a
}

// @lc code=end

