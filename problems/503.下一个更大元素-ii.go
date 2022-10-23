/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	stack := make([]int, 0)
	nlen := len(nums)

	for i := 2*nlen - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%nlen] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%nlen] = -1
		} else {
			res[i%nlen] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%nlen])
	}

	return res
}

// @lc code=end

