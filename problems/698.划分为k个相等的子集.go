/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
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

// @lc code=end

