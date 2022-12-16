/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var dfs func(left int, path []int)
	dfs = func(left int, path []int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}

		for i := left; i <= n; i++ {
			path = append(path, i)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(1, []int{})
	return res
}

// @lc code=end

