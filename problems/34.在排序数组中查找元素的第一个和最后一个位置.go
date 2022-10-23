/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := leftBound(nums, target, function)
	right := rightBound(nums, target, function)
	return []int{left, right}
}

type F func(nums []int, x int) int

func function(nums []int, x int) int {
	return nums[x]
}

func leftBound(nums []int, target int, f F) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if f(nums, mid) == target {
			right = mid
		} else if f(nums, mid) < target {
			left = mid + 1
		} else if f(nums, mid) > target {
			right = mid
		}
	}

	if f(nums, left) == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int, f F) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2 + 1
		if f(nums, mid) == target {
			left = mid
		} else if f(nums, mid) < target {
			left = mid
		} else if f(nums, mid) > target {
			right = mid - 1
		}
	}

	if f(nums, left) == target {
		return left
	}
	return -1
}

// @lc code=end

