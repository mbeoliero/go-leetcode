/*
 * @lc app=leetcode.cn id=969 lang=golang
 *
 * [969] 煎饼排序
 */

// @lc code=start
var res []int

func pancakeSort(arr []int) []int {
	res = make([]int, 0)
	sort(arr, len(arr))
	return res
}

func sort(arr []int, n int) {
	if n == 1 {
		return
	}

	var max, maxIndex int
	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
			maxIndex = i
		}
	}

	reverse(arr, 0, maxIndex)
	res = append(res, maxIndex+1)
	reverse(arr, 0, n-1)
	res = append(res, n)
	sort(arr, n-1)
}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// @lc code=end

