/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	var res int = 4

	var dfs func(num int, depth int, last int)
	dfs = func(num int, depth int, last int) {
		if num == 0 {
			res = min(res, depth)
			return
		}
		if res <= 2 || depth >= 3 {
			return
		}

		up := min(last, int(math.Sqrt(float64(num))))
		low := max(1, int(math.Sqrt(float64(num/(3-depth)))))
		for i := up; i >= low; i-- {
			dfs(num-i*i, depth+1, i)
		}
		return
	}

	dfs(n, 0, n)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

