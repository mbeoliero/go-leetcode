/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	count := make([][]int, n)
	var res int

	for i := 0; i < n; i++ {
		count[i] = make([]int, m)

		sum := 0
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				sum++
				count[i][j] = sum
				if count[i][j] > res {
					res = count[i][j]
				}
			} else {
				sum = 0
			}
		}
	}

	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if count[i][j] > 0 {
				width := count[i][j]
				for k := i - 1; k >= 0; k-- {
					if width > count[k][j] {
						width = count[k][j]
					}
					temp := (i - k + 1) * width
					if temp > res {
						res = temp
					}
				}
			}
		}
	}

	return res
}

// @lc code=end

