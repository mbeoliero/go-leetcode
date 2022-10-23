/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start

var res [][]int
var sum int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	path := make([]int, 0)
	sort.IntSlice(candidates).Sort()
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
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		sum += nums[i]
		backTrack(nums, i+1, target, path)
		path = path[:len(path)-1]
		sum -= nums[i]
	}
}

// @lc code=end

