/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	piv := -1
	p := len(nums) - 1
	for p > 0 {
		if nums[p-1] < nums[p] {
			piv = p - 1
			break
		}
		p--
	}
	if piv == -1 {
		reverse(&nums, 0, len(nums)-1)
		return
	}

	nextPiv := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[piv] {
			nextPiv = i
			break
		}
	}

	nums[piv], nums[nextPiv] = nums[nextPiv], nums[piv]
	reverse(&nums, piv+1, len(nums)-1)
	return
}

func reverse(nums *[]int, start, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}

// @lc code=end

