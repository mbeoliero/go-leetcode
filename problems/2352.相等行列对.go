/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] 相等行列对
 */

// @lc code=start
func equalPairs(grid [][]int) (ans int) {
	var rowmap = make(map[string]int)
	var colmap = make(map[string]int)
	var count int
	n := len(grid)
	for i := 0; i < n; i++ {
		var row string
		for j := 0; j < n; j++ {
			row += strconv.Itoa(grid[i][j])
			row += ","
		}
		rowmap[row]++
	}
	for i := 0; i < n; i++ {
		var col string
		for j := 0; j < n; j++ {
			col += strconv.Itoa(grid[j][i])
			col += ","
		}
		colmap[col]++
		if rowmap[col] >= 1 {
			count += rowmap[col]
		}
	}
	return count
}

// @lc code=end

