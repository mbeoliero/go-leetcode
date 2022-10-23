/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	i := 0
	for _, v := range nums {
		if v != 0 {
			nums[i] = v
			i++
		}
	}
	for ; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

