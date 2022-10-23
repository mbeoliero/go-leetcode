/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
var count []int
var temp []Pair

type Pair struct {
	Val int
	Idx int
}

func countSmaller(nums []int) []int {
	temp = make([]Pair, len(nums))
	count = make([]int, len(nums))
	n := make([]Pair, len(nums))
	for i, nu := range nums {
		n[i] = Pair{Val: nu, Idx: i}
	}

	mergeSort(n, 0, len(n)-1)
	return count
}

func mergeSort(nums []Pair, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2
	mergeSort(nums, low, mid)
	mergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []Pair, low, mid, high int) {
	for i := low; i <= high; i++ {
		temp[i] = nums[i]
	}

	i, j := low, mid+1
	p := low
	for i <= mid && j <= high {
		if temp[i].Val <= temp[j].Val {
			nums[p] = temp[i]
			count[temp[i].Idx] += j - mid - 1
			i++
			p++
		} else {
			nums[p] = temp[j]
			j++
			p++

		}
	}

	for i <= mid {
		nums[p] = temp[i]
		count[temp[i].Idx] += high - mid
		i++
		p++
	}
	for j <= high {
		nums[p] = temp[j]
		j++
		p++
	}
}

// @lc code=end

