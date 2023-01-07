/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}

	ans := len(nums) + 1
	left, right := -1, 0
	lsum, rsum := 0, sum
	for ; left < len(nums); left++ {
		if left != -1 {
			lsum += nums[left]
		}

		for right < len(nums) && lsum+rsum > x {
			rsum -= nums[right]
			right++
		}

		if lsum+rsum == x {
			if ans > left+1+len(nums)-right {
				ans = left + 1 + len(nums) - right
			}
		}
	}
	if ans > len(nums) {
		return -1
	}
	return ans
}

// @lc code=end

