/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	res := make([]int, length)
	for _, v := range nums {
		res[v-1] = v
	}
	j := 0
	for i := 0; i < length; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

// @lc code=end

