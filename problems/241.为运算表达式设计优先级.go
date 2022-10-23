/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	res := make([]int, 0)

	for i := 0; i < len(expression); i++ {
		if expression[i] == '-' || expression[i] == '+' || expression[i] == '*' {
			// 分
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			// 治
			for _, l := range left {
				for _, r := range right {
					switch expression[i] {
					case '-':
						res = append(res, l-r)
					case '+':
						res = append(res, l+r)
					case '*':
						res = append(res, l*r)
					}
				}
			}
		}
	}

	// 递归的出口
	if len(res) == 0 {
		in, _ := strconv.Atoi(expression)
		res = append(res, in)
	}
	return res
}

// @lc code=end

