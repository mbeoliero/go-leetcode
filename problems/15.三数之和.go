/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
var res [][]int

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.IntSlice(nums).Sort()
	for i := 0; i < len(nums); i++ {
		two := twoSum(nums[i+1:], 0-nums[i])
		if len(two) > 0 {
			for _, v := range two {
				res = append(res, append(v, nums[i]))
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSum(numbers []int, target int) [][]int {
	res := make([][]int, 0)
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			res = append(res, []int{numbers[left], numbers[right]})
			for left < len(numbers)-1 && numbers[left] == numbers[left+1] {
				left++
			}
			for right > 1 && numbers[right] == numbers[right-1] {
				right--
			}
		}
		if numbers[left]+numbers[right] < target {
			for left < len(numbers)-1 && numbers[left] == numbers[left+1] {
				left++
			}
			left++
		} else {
			for right > 1 && numbers[right] == numbers[right-1] {
				right--
			}
			right--
		}
	}
	return res
}

// @lc code=end

