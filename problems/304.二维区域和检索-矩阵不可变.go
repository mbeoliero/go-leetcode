/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	SumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sum[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		temp := 0
		for j := 0; j < n; j++ {
			temp += matrix[i][j]
			sum[i+1][j+1] = sum[i][j+1] + temp
		}
	}
	return NumMatrix{
		SumMatrix: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.SumMatrix[row2+1][col2+1] - this.SumMatrix[row1][col2+1] - this.SumMatrix[row2+1][col1] + this.SumMatrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

