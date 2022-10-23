/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */
// @lc code=start
import "math"

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

// @lc code=end
