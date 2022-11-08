/*
 * @lc app=leetcode.cn id=1678 lang=golang
 *
 * [1678] 设计 Goal 解析器
 */

// @lc code=start
// func interpret(command string) string {
// 	res := ""

// 	i := 0
// 	for i < len(command) {
// 		switch command[i] {
// 		case 'G':
// 			res += "G"
// 			i++
// 		case '(':
// 			if i+1 < len(command) && command[i+1] == ')' {
// 				res += "o"
// 				i = i + 2
// 			}
// 			if i+4 <= len(command) && command[i:i+4] == "(al)" {
// 				res += "al"
// 				i = i + 4
// 			}
// 		default:
// 			i++
// 		}
// 	}

// 	return res
// }
func interpret(command string) string {
	res := &strings.Builder{}
	for i, c := range command {
		if c == 'G' {
			res.WriteByte('G')
		} else if c == '(' {
			if command[i+1] == ')' {
				res.WriteByte('o')
			} else {
				res.WriteString("al")
			}
		}
	}
	return res.String()
}

// @lc code=end

