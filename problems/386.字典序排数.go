/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, 0)

	var dfs func(cur, n int)
	dfs = func(cur, n int) {
		if cur <= n {
			res = append(res, cur)
			for i := 0; i <= 9; i++ {
				if cur*10+i <= n {
					dfs(cur*10+i, n)
				}
			}
		}
	}

	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}
	return res
}

// @lc code=end

