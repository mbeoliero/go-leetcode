/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, n := range nums {
		if left < n {
			left = n
		}
		right += n
	}

	for left < right {
		mid := left + (right-left)/2
		if function(nums, mid) == k {
			right = mid
		} else if function(nums, mid) < k {
			right = mid
		} else if function(nums, mid) > k {
			left = mid + 1
		}
	}
	return left
}

func function(nums []int, sum int) int {
	res := 0
	temp := 0
	for _, n := range nums {
		if temp+n > sum {
			res++
			temp = 0
		}
		temp += n
	}
	return res + 1
}

// @lc code=end

