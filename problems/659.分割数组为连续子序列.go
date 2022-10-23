/*
 * @lc app=leetcode.cn id=659 lang=golang
 *
 * [659] 分割数组为连续子序列
 */

// @lc code=start
func isPossible(nums []int) bool {
	freq, need := make(map[int]int), make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}

		if need[v] > 0 {
			freq[v]--
			need[v]--
			need[v+1]++
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 {
			freq[v]--
			freq[v+1]--
			freq[v+2]--
			need[v+3]++
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

