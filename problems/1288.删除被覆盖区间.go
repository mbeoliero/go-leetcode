/*
 * @lc app=leetcode.cn id=1288 lang=golang
 *
 * [1288] 删除被覆盖区间
 */

// @lc code=start
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][1] >= intervals[i+1][1] {
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
		}
	}

	return len(intervals)
}

// @lc code=end

