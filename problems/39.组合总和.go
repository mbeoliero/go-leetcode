/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
var res [][]int
var sum int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	path := make([]int, 0)
	backTrack(candidates, 0, target, path)
	return res
}

func backTrack(nums []int, start, target int, path []int) {
	if sum == target {
		res = append(res, append([]int{}, path...))
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		sum += nums[i]
		backTrack(nums, i, target, path)
		path = path[:len(path)-1]
		sum -= nums[i]
	}
}

// @lc code=end

