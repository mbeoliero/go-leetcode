/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := m-1, 0
	for i >= 0 && j < n {
		for i >= 0 {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] > target {
				i--
			} else {
				break
			}
		}

		if i < 0 {
			break
		}

		for j < n {
			if matrix[i][j] == target {
				return true
			} else if matrix[i][j] < target {
				j++
			} else {
				break
			}
		}
	}

	return false
}

// @lc code=end

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1

	for i < n && j >= 0 {
		v := matrix[i][j]
		if v == target {
			return true
		} else if v > target {
			j -= 1
		} else {
			i += 1
		}

	}
	return false
}