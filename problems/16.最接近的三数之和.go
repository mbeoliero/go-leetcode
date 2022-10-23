/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt16

	update := func(cur int) {
		if abs(cur-target) < abs(res-target) {
			res = cur
		}
	}

	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == target {
				return target
			}

			update(sum)
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

