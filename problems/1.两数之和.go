package problems

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, val := range nums {
		if i, ok := m[val]; ok {
			return []int{i, idx}
		}
		m[target-val] = idx
	}
	return []int{}
}

// @lc code=end
