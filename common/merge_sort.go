package common

// func MergeSort(nums []int) []int {
// 	if len(nums) < 2 {
// 		return nums
// 	}
// 	m := len(nums) / 2
// 	l := MergeSort(nums[:m])
// 	r := MergeSort(nums[m:])
// 	return merge(l, r)
// }

// func merge(nums1, nums2 []int) (res []int) {
// 	i, j := 0, 0
// 	for i < len(nums1) && j < len(nums2) {
// 		if nums1[i] <= nums2[j] {
// 			res = append(res, nums1[i])
// 			i++
// 		} else {
// 			res = append(res, nums2[j])
// 			j++
// 		}
// 	}
// 	res = append(res, nums1[i:]...)
// 	res = append(res, nums2[j:]...)
// 	return
// }

var temp []int

func MergeSort(nums []int) {
	temp = make([]int, len(nums))
	ParQuickSort(nums, 0, len(nums)-1)
}

func ParMergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	ParMergeSort(nums, low, mid)
	ParMergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		temp[i] = nums[i]
	}

	i, j := low, mid+1
	p := low
	for i <= mid && j <= high {
		if temp[i] <= temp[j] {
			nums[p] = temp[i]
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
		i++
		p++
	}
	for j <= high {
		nums[p] = temp[j]
		j++
		p++
	}
}
