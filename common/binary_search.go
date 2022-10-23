package common

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

type F func(nums []int, x int) int

func function(nums []int, x int) int {
	return nums[x]
}

func LeftBoundSearch(nums []int, target int) int {
	return leftBound(nums, target, function)
}

func RightBoundSearch(nums []int, target int) int {
	return rightBound(nums, target, function)
}

func leftBound(nums []int, target int, f F) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if f(nums, mid) == target {
			right = mid
		} else if f(nums, mid) < target {
			left = mid + 1
		} else if f(nums, mid) > target {
			right = mid
		}
	}

	if f(nums, left) == target {
		return left
	}
	return -1
}

func rightBound(nums []int, target int, f F) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2 + 1
		if f(nums, mid) == target {
			left = mid
		} else if f(nums, mid) < target {
			left = mid
		} else if f(nums, mid) > target {
			right = mid - 1
		}
	}

	if f(nums, left) == target {
		return left
	}
	return -1
}
