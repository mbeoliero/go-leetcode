/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	h := make([]int, 0, len(envelopes))
	for _, env := range envelopes {
		h = append(h, env[1])
	}

	return lengthOfLIS(h)
}

func lengthOfLIS(nums []int) int {
	tail := make([]int, 0)
	for _, n := range nums {
		idx := binarySearch(tail, n)
		if idx == len(tail) {
			tail = append(tail, n)
		} else {
			tail[idx] = n
		}
	}

	return len(tail)
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

// @lc code=end

