/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// var res [][]int

// func permute(nums []int) [][]int {
// 	res = [][]int{}
// 	backTrack(nums, len(nums), []int{})
// 	return res
// }

// func backTrack(nums []int, numsLen int, path []int) {
// 	if len(nums) == 0 {
// 		p := make([]int, len(path))
// 		copy(p, path)
// 		res = append(res, p)
// 	}
// 	for i := 0; i < numsLen; i++ {
// 		cur := nums[i]
// 		path = append(path, cur)
// 		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
// 		backTrack(nums, len(nums), path)
// 		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
// 		path = path[:len(path)-1]
// 	}
// }

func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

// We use a pointer for the result so we don't need to worry returning it.
func permuteRec(currComb, left []int, res *[][]int) {
	// We know that we found a new combination when we have no elements left.
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	// For the next iteration we consider all the left elements but the current one (idx).
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...), // Make sure to allocate a new slice.
			res,
		)
	}
}

// @lc code=end

func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

// We use a pointer for the result so we don't need to worry returning it.
func permuteRec(currComb, left []int, res *[][]int) {
	// We know that we found a new combination when we have no elements left.
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	// For the next iteration we consider all the left elements but the current one (idx).
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...), // Make sure to allocate a new slice.
			res,
		)
	}
}