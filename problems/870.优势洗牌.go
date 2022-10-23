/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	b := make([][2]int, len(nums2))
	for i, n := range nums2 {
		b[i] = [2]int{n, i}
	}

	sort.Slice(b, func(i, j int) bool {
		return b[i][0] > b[j][0]
	})

	left, right := 0, len(nums1)-1
	res := make([]int, len(nums1))
	for _, v := range b {
		if v[0] >= nums1[right] {
			res[v[1]] = nums1[left]
			left++
		} else {
			res[v[1]] = nums1[right]
			right--
		}
	}
	return res
}

// @lc code=end

