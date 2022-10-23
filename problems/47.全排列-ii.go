/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	sort.IntSlice(nums).Sort()
	backTrack(nums, len(nums), []int{})
	return res
}

func backTrack(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]
	}
}

// @lc code=end

