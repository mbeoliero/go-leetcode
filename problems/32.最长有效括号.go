/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
// Support function max
func Max(a, b int) int {

	if a > b {
		return a
	}

	return b
}

func longestValidParentheses(s string) int {

	// stack, used to record index of parenthesis
	// initialized to -1 as dummy head for valid parenthesis length computation
	stack := []int{-1}

	maxLength := 0

	// linear scan each index and character in input string s
	for curIdx, char := range s {

		if char == '(' {

			// push when current char is (
			stack = append(stack, curIdx)

		} else {

			// pop when current char is )
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {

				// stack is empty, push current index into stack
				stack = append(stack, curIdx)
			} else {

				// stack is non-empty, update maximal valud parentheses length

				maxLength = Max(maxLength, curIdx-stack[len(stack)-1])
			}

		}

	}

	return maxLength

}

// ()(())
// 0  -1 0
// 2  -1
// 2  -1 2
// 2  -1 2 3
// 2  -1 2
// 6  -1

// @lc code=end

