/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	tail := make([]int, 0)
	for _, n := range nums {
		idx := binarySearch(tail, n)
		if idx == len(tail) {
			tail = append(tail, n)
		} else {
			tail[idx] = n
		}
	}

	return len(tail)
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

// func lengthOfLIS(nums []int) int {
// 	var res int
// 	dp := make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		dp[i] = 1
// 		for j := i; j >= 0; j-- {
// 			if nums[i] > nums[j] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		res = max(res, dp[i])
// 	}
// 	return res
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

