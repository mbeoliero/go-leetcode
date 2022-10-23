/*
 * @lc app=leetcode.cn id=1081 lang=golang
 *
 * [1081] 不同字符的最小子序列
 */

// @lc code=start
func smallestSubsequence(s string) string {
	count := [256]int{}
	instack := [256]bool{}
	stack := make([]rune, 0)

	for _, r := range s {
		count[r]++
	}

	for _, r := range s {
		count[r]--
		if instack[r] {
			continue
		}

		for {
			last := len(stack) - 1
			if last >= 0 && stack[last] > r && count[stack[last]] > 0 {
				instack[stack[last]] = false
				stack = stack[:last]
			} else {
				break
			}
		}

		stack = append(stack, r)
		instack[r] = true
	}
	return string(stack)
}

// @lc code=end

