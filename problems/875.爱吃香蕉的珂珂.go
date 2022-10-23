/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, p := range piles {
		if right < p {
			right = p
		}
	}

	for left < right {
		mid := left + (right-left)/2
		if function(piles, mid) == h {
			right = mid
		} else if function(piles, mid) < h {
			right = mid
		} else if function(piles, mid) > h {
			left = mid + 1
		}
	}
	return left
}

func function(piles []int, k int) int {
	var h int
	for _, p := range piles {
		if p%k == 0 {
			h += p / k
		} else {
			h += p/k + 1
		}
	}
	return h
}

// @lc code=end

