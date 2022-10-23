/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || 0 == k {
		return make([]int, 0)
	} else if 1 == k {
		return nums
	}

	var (
		// len(nums)-k+1 this is the number of windows.
		// If input has length 3 and k is 1, we'll have 3 - 1 + 1 = 3 windows.
		res = make([]int, len(nums)-k+1)
		dq  = &deque{}
	)

	for i := range nums {
		// we pop the first element if it's outside of the current window.
		if false == dq.empty() && (i-k == dq.getFirst()) {
			dq.popFirst()
		}

		// we pop all the elements that are smaller than the current one.
		for false == dq.empty() && nums[dq.getLast()] < nums[i] {
			dq.popLast()
		}

		// we push the current one to the window.
		dq.push(i)

		// if we reached at least the first window end.
		if i >= k-1 {
			// we add the current result that is the first in the deque.
			res[i-k+1] = nums[dq.getFirst()]
		}
	}

	return res
}

type deque struct {
	indexes []int
}

func (d *deque) push(i int) {
	d.indexes = append(d.indexes, i)
}

func (d *deque) getFirst() int {
	return d.indexes[0]
}

func (d *deque) popFirst() {
	d.indexes = d.indexes[1:]
}

func (d *deque) getLast() int {
	return d.indexes[len(d.indexes)-1]
}

func (d *deque) popLast() {
	d.indexes = d.indexes[:len(d.indexes)-1]
}

func (d *deque) empty() bool {
	return 0 == len(d.indexes)
}

// @lc code=end

