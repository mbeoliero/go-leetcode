/*
 * @lc app=leetcode.cn id=805 lang=golang
 *
 * [805] 数组的均值分割
 */

// @lc code=start

func splitArraySameAverage(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}

	n := len(nums)
	m := n / 2
	isPossible := false

	for i := 1; i <= m; i++ {
		if sum*i%n == 0 {
			isPossible = true
			break
		}
	}

	if !isPossible {
		return false
	}

	dp := make([]map[int]bool, m+1)
	for i := range dp {
		dp[i] = map[int]bool{}
	}

	dp[0][0] = true
	for _, num := range nums {
		for i := m; i >= 1; i-- {
			for x := range dp[i-1] {
				curr := x + num
				if curr*n == sum*i {
					return true
				}
				dp[i][curr] = true
			}
		}
	}

	return false
}

// @lc code=end

func splitArraySameAverage(nums []int) bool {
	nlen := len(nums)
	if nlen == 1 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}

	for i := range nums {
		// 避免引入浮点数
		nums[i] = nums[i]*nlen - sum
	}

	m := nlen / 2
	left := map[int]bool{}

	for i := 1; i < 1<<m; i++ {
		total := 0
		for j, x := range nums[:m] {
			if i>>j&1 > 0 {
				total += x
			}
		}

		if total == 0 {
			return true
		}
		left[total] = true
	}

	right := 0
	for _, x := range nums[m:] {
		right += x
	}

	for i := 1; i < 1<<(nlen-m); i++ {
		total := 0
		for j, x := range nums[m:] {
			if i>>j&1 > 0 {
				total += x
			}
		}

		if total == 0 || right != total && left[-total] {
			return true
		}
	}

	return false
}
