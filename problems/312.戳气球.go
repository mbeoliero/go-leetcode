/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
var res int

func maxCoins(nums []int) int {
	nlen := len(nums)
	dp := make([][]int, nlen+2)
	for i := range dp {
		dp[i] = make([]int, nlen+2)
	}

	points := make([]int, nlen+2)
	points[0], points[nlen+1] = 1, 1
	for i := 1; i < nlen+1; i++ {
		points[i] = nums[i-1]
	}

	// dp[i][j]表示（i，j）开区间内能获得最多的金币数
	for i := nlen; i >= 0; i-- {
		for j := i + 1; j < nlen+2; j++ {
			for k := i + 1; k < j; k++ {
				// k是剩到最后戳的气球
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[k]*points[j])
			}
		}
	}

	return dp[0][nlen+1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

func maxCoins(nums []int) int {
	n := len(nums)
	val := make([]int, n+2)
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	val[0], val[n+1] = 1, 1
	rec := make([][]int, n+2)
	for i := 0; i < len(rec); i++ {
		rec[i] = make([]int, n+2)
		for j := 0; j < len(rec[i]); j++ {
			rec[i][j] = -1
		}
	}
	return solve(0, n+1, val, rec)
}

func solve(left, right int, val []int, rec [][]int) int {
	if left >= right-1 {
		return 0
	}
	if rec[left][right] != -1 {
		return rec[left][right]
	}
	for i := left + 1; i < right; i++ {
		sum := val[left] * val[i] * val[right]
		sum += solve(left, i, val, rec) + solve(i, right, val, rec)
		rec[left][right] = max(rec[left][right], sum)
	}
	return rec[left][right]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}