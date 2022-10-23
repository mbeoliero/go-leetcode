/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum < target || (target < 0 && -target > sum) || (sum+target)%2 == 1 {
		return 0
	}

	sum = (sum + target) / 2
	return subsets(nums, sum)
}

func subsets(nums []int, sum int) int {
	n := len(nums)
	// dp := make([][]int, n+1)
	// for i := 0; i < n+1; i++ {
	// 	dp[i] = make([]int, sum+1)
	// }
	dp := make([]int, sum+1)
	// dp[0][0] = 1
	dp[0] = 1

	// for i := 1; i < n+1; i++ {
	// 	for j := 0; j < sum+1; j++ {
	// 		if j < nums[i-1] {
	// 			dp[i][j] = dp[i-1][j]
	// 		} else {
	// 			dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
	// 		}
	// 	}
	// }
	for i := 1; i < n+1; i++ {
		for j := sum; j >= 0; j-- {
			if j < nums[i-1] {
				dp[j] = dp[j]
			} else {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			}
		}
	}

	// return dp[n][sum]
	return dp[sum]
}

// @lc code=end

var result int
var tar int
var op []int

func findTargetSumWays(nums []int, target int) int {
	op = []int{1, -1}
	tar = target
	result = 0
	backtrack(nums, 0, 0)
	return result
}

func backtrack(nums []int, sum, index int) {
	if index == len(nums) {
		if sum == tar {
			result++
		}
		return
	}

	for _, o := range op {
		sum += o * nums[index]
		backtrack(nums, sum, index+1)
		sum -= o * nums[index]
	}
}
