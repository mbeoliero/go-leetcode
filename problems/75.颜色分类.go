/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	const RED, WHITE, BLUE = 0, 1, 2

	i, idx_red, idx_blue := 0, 0, len(nums)-1

	for ; i <= idx_blue; i += 1 {

		if nums[i] == RED {
			nums[i], nums[idx_red] = nums[idx_red], nums[i]
			idx_red += 1

		} else if nums[i] == BLUE {
			nums[i], nums[idx_blue] = nums[idx_blue], nums[i]
			idx_blue -= 1

			// i stay here for one more check on next iteration
			i -= 1

		}

	}

	return
}

// @lc code=end

