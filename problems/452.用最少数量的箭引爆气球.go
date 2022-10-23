/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1 //计算不相交的区间数
	end := points[0][1]

	for i := 0; i < len(points); i++ {
		if points[i][0] > end {
			count++
			end = points[i][1]
		}
	}

	return count
}

// @lc code=end

