/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[nums2[i]] = -1
		} else {
			res[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	res1 := make([]int, len(nums1))
	for i, n := range nums1 {
		res1[i] = res[n]
	}
	return res1
}

// @lc code=end

