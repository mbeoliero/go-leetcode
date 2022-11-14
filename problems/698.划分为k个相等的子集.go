/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start

func canPartitionKSubsets(nums []int, k int) bool {
	all := 0
	for _, v := range nums {
		all += v
	}
	if all%k > 0 {
		return false
	}
	per := all / k
	sort.Ints(nums)
	n := len(nums)
	if nums[n-1] > per {
		return false
	}

	dp := make([]bool, 1<<n)
	for i := range dp {
		dp[i] = true
	}
	var dfs func(int, int) bool
	dfs = func(s, p int) bool {
		if s == 0 {
			return true
		}
		if !dp[s] {
			return dp[s]
		}
		dp[s] = false
		for i, num := range nums {
			if num+p > per {
				break
			}
			if s>>i&1 > 0 && dfs(s^1<<i, (p+nums[i])%per) {
				return true
			}
		}
		return false
	}
	return dfs(1<<n-1, 0)
}

// @lc code=end

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	bucket := make([]int, k)
	return backtrack(nums, bucket, target, 0)
}

func backtrack(nums []int, bucket []int, target int, index int) bool {
	if index >= len(nums) {
		// for i := 0; i < len(bucket); i++ {
		// 	if bucket[i] != target {
		// 		return false
		// 	}
		// }
		return true
	}

	for i := 0; i < len(bucket); i++ {
		if i > 0 && bucket[i] == bucket[i-1] {
			continue
		}
		if bucket[i]+nums[index] > target || bucket[i] == target {
			continue
		}

		bucket[i] += nums[index]
		if backtrack(nums, bucket, target, index+1) {
			return true
		}

		bucket[i] -= nums[index]
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	n := len(nums)
	if nums[n-1] > target {
		return false
	}

	mask := (1 << n) - 1
	f := make([]int, 1<<n)

	var dfs func(state int, per int) bool
	dfs = func(state int, per int) bool {
		if state == mask {
			return true
		}
		if f[state] != 0 {
			return f[state] == 1
		}

		for i, num := range nums {
			// 第 i 个元素已添加
			if (state >> i & 1) == 1 {
				continue
			}

			if num+per > target {
				break
			}

			if dfs(state|1<<i, (per+num)%target) {
				f[state] = 1
				return true
			}
		}

		f[state] = -1
		return false
	}

	return dfs(0, 0)
}