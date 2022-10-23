/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
var res [][]int

func subsets(nums []int) [][]int {
	res = [][]int{}
	path := make([]int, 0)
	backTrack(nums, 0, path)
	return res
}

func backTrack(nums []int, start int, path []int) {
	res = append(res, append([]int{}, path...))
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backTrack(nums, i+1, path)
		path = path[:len(path)-1]
	}
}

// func subsets(nums []int) [][]int {
// 	ans := make([][]int, 0)
// 	track := make([]int, 0)
// 	var backtrack func(start int)
// 	backtrack = func(start int) {
// 		ans = append(ans, append([]int{}, track...))
// 		for i := start; i < len(nums); i++ {
// 			track = append(track, nums[i])
// 			backtrack(i + 1)
// 			track = track[:len(track)-1]
// 		}
// 	}

// 	backtrack(0)
// 	return ans
// }

// @lc code=end

