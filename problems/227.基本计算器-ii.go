/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	stack := make([]int, 0)
	sign := '+'
	num := 0

	for i := 0; i < len(s); i++ {

		if isdigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}

		if !isdigit(s[i]) && s[i] != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre*num)
			case '/':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre/num)
			}

			sign = rune(s[i])
			num = 0
		}
	}

	res := 0
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i]
	}
	return res
}

func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// @lc code=end

