/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = [][]int{}
	path := make([]int, 0)
	sort.IntSlice(nums).Sort()
	backTrack(nums, 0, path)
	return res
}

func backTrack(nums []int, start int, path []int) {
	res = append(res, append([]int{}, path...))
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backTrack(nums, i+1, path)
		path = path[:len(path)-1]
	}
}

// @lc code=end

