/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

type heapList []int

func (h heapList) Len() int {
	return len(h)
}

func (h heapList) Less(i, j int) bool {
	return h[i] <= h[j]
}

func (h heapList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapList) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h *heapList) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func findKthLargest(nums []int, k int) int {
	hp := &heapList{}
	for i, n := range nums {
		if i < k {
			heap.Push(hp, n)
		} else {
			heap.Push(hp, n)
			heap.Pop(hp)
		}
	}

	return heap.Pop(hp).(int)
}

// func findKthLargest(nums []int, k int) int {
// 	nlen := len(nums)
// 	target := nlen - k
// 	left, right := 0, nlen-1

// 	for {
// 		p := partition(nums, left, right)
// 		if p == target {
// 			return nums[p]
// 		}
// 		if p < target {
// 			left = p + 1
// 		} else {
// 			right = p - 1
// 		}
// 	}
// 	return 0
// }

// func partition(list []int, low, high int) int {
// 	pivot := list[low]
// 	for low < high {
// 		for low < high && pivot <= list[high] {
// 			high--
// 		}
// 		list[low] = list[high]

// 		for low < high && pivot > list[low] {
// 			low++
// 		}
// 		list[high] = list[low]
// 	}

// 	list[low] = pivot
// 	return low
// }

// @lc code=end

