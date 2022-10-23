/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	k := 1
	left, right := 0, n-1
	upper, lower := 0, n-1

	for k <= n*n {
		if upper <= lower {
			for j := left; j <= right; j++ {
				res[upper][j] = k
				k++
			}
			upper++
		}

		if left <= right {
			for i := upper; i <= lower; i++ {
				res[i][right] = k
				k++
			}
			right--
		}

		if upper <= lower {
			for j := right; j >= left; j-- {
				res[lower][j] = k
				k++
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				res[i][left] = k
				k++
			}
			left++
		}
	}
	return res
}

// @lc code=end

