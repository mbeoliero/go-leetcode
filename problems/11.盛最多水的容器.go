/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	var res int
	left, right := 0, len(height)-1
	for left < right {
		temp := min(height[left], height[right]) * (right - left)
		if res < temp {
			res = temp
		}

		if min(height[left], height[right]) == height[right] {
			i := right - 1
			for ; i >= left && height[i] <= height[right]; i-- {
			}
			right = i
		} else {
			i := left + 1
			for ; i <= right && height[i] <= height[left]; i++ {
			}
			left = i
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func Max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

func maxArea(height []int) int {

	// length of input array
	size := len(height)

	// tow pointers, left init as 0, right init as size-1
	left, right := 0, size-1

	// maximal width between leftmost stick and rightmost stick
	maxWidth := size - 1

	// area also known as the amount of water
	area := 0

	// trade-off between width and height
	// scan each possible width and compute maximal area
	for width := maxWidth; width > 0; width-- {

		if height[left] < height[right] {

			// the height of lefthand side is shorter
			area = Max(area, width*height[left])

			// update left index to righthand side
			left += 1

		} else {

			// the height of righthand side is shorter
			area = Max(area, width*height[right])

			// update left index to righthand side
			right -= 1
		}

	}
	return area
}