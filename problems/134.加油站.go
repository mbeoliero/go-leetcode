/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum, minSum := 0, 0
	start := 0

	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			start = i + 1
			minSum = sum
		}
	}

	if sum < 0 {
		return -1
	}
	if start == n {
		start = 0
	}
	return start
}

// @lc code=end

