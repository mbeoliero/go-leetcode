/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// func fib(n int) int {
// 	if n == 0 || n == 1 {
// 		return n
// 	}
// 	dp := make([]int, n+1)
// 	dp[0], dp[1] = 0, 1
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1] + dp[i-2]
// 	}
// 	return dp[n]
// }

func fib(n int) int {
	var a, b, c int
	a, b = 0, 1
	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return a
}

// @lc code=end

