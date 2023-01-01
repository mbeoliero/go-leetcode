/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}

	target := sum / 4

	n := len(matchsticks)
	dp := make([]int, 1<<n)
	var dfs func(state int, per int) bool
	dfs = func(state int, per int) bool {
		if state == (1<<n)-1 {
			return true
		}
		if dp[state] != 0 {
			return dp[state] == 1
		}

		for i, num := range matchsticks {
			if state&(1<<i) > 0 {
				continue
			}
			if num+per > target {
				break
			}

			if dfs(state|(1<<i), (per+num)%target) {
				dp[state] = 1
				return true
			}
		}
		dp[state] = -1
		return false
	}
	return dfs(0, 0)
}

// @lc code=end

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	target := sum / 4
	buckets := make([]int, 4)

	var backtrack func(nums []int, buckets []int, target, index int) bool
	backtrack = func(nums []int, buckets []int, target, index int) bool {
		if index == len(nums) {
			return true
		}

		for i := 0; i < len(buckets); i++ {
			if buckets[i]+nums[index] > target {
				continue
			}

			buckets[i] += nums[index]
			if backtrack(nums, buckets, target, index+1) {
				return true
			}
			buckets[i] -= nums[index]
		}
		return false
	}

	return backtrack(matchsticks, buckets, target, 0)
}