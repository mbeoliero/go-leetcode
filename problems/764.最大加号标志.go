/*
 * @lc app=leetcode.cn id=764 lang=golang
 *
 * [764] 最大加号标志
 */

// @lc code=start

func orderOfLargestPlusSign(n int, mines [][]int) (ans int) {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = n
		}
	}
	banned := map[int]bool{}
	for _, p := range mines {
		banned[p[0]*n+p[1]] = true
	}
	for i := 0; i < n; i++ {
		count := 0
		/* left */
		for j := 0; j < n; j++ {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
		count = 0
		/* right */
		for j := n - 1; j >= 0; j-- {
			if banned[i*n+j] {
				count = 0
			} else {
				count++
			}
			dp[i][j] = min(dp[i][j], count)
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		/* up */
		for j := 0; j < n; j++ {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
		}
		count = 0
		/* down */
		for j := n - 1; j >= 0; j-- {
			if banned[j*n+i] {
				count = 0
			} else {
				count++
			}
			dp[j][i] = min(dp[j][i], count)
			ans = max(ans, dp[j][i])
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// @lc code=end

func orderOfLargestPlusSign(n int, mines [][]int) int {
	if len(mines) == n*n {
		return 0
	}

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, m := range mines {
		grid[m[0]][m[1]] = 1
	}

	for next := 0; next < (n+1)/2-1; next++ {
		points := getCenter(n, next)
		for _, p := range points {
			if w, b := getPlusSign(n, grid, p[0], p[1]); b {
				return w
			}
		}
	}

	return 1
}

func getCenter(n int, next int) [][]int {
	res := make([][]int, 0)
	center := n / 2

	if n%2 == 1 {
		for i := center - next; i <= center+next; i++ {
			res = append(res, []int{i, center - next})
			res = append(res, []int{i, center + next})
		}

		for j := center - next + 1; j < center+next; j++ {
			res = append(res, []int{center - next, j})
			res = append(res, []int{center + next, j})
		}
		return res
	}

	for i := center - 1 - next; i <= center+next; i++ {
		res = append(res, []int{i, center - 1 - next})
		res = append(res, []int{i, center + next})
	}

	for j := center - next; j < center+next; j++ {
		res = append(res, []int{center - 1 - next, j})
		res = append(res, []int{center + next, j})
	}
	return res
}

func getPlusSign(n int, grid [][]int, i, j int) (int, bool) {
	width := min(i-0, n-i, j-0, n-j)

	for p := i - width; p <= i+width; p++ {
		if grid[p][j] == 1 {
			return 0, false
		}
	}

	for p := j - width; p <= j+width; p++ {
		if grid[i][p] == 1 {
			return 0, false
		}
	}

	return width + 1, true
}

func min(val ...int) int {
	m := 1000
	for _, v := range val {
		if m > v {
			m = v
		}
	}
	return m
}