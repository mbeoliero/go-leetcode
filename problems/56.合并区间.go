/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] && intervals[i][1] > res[len(res)-1][1] {
			res[len(res)-1][1] = intervals[i][1]
		}
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		}
	}
	return res
}

// @lc code=end

