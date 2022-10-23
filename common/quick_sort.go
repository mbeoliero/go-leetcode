package common

func QuickSort(nums []int) {
	ParQuickSort(nums, 0, len(nums)-1)
}

func ParQuickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	ParQuickSort(nums, low, p-1)
	ParQuickSort(nums, p+1, high)
}

func partition(list []int, low, high int) int {
	pivot := list[low]
	for low < high {
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]

		for low < high && pivot > list[low] {
			low++
		}
		list[high] = list[low]
	}

	list[low] = pivot
	return low
}
