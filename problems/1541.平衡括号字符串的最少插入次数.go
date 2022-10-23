/*
 * @lc app=leetcode.cn id=1541 lang=golang
 *
 * [1541] 平衡括号字符串的最少插入次数
 */

// @lc code=start
func minInsertions(s string) int {
	counter := 0
	insertion := 0
	for _, r := range s {
		if r == '(' {
			if counter%2 == 1 {
				// insert a ')'
				insertion++
				counter--
			}
			counter += 2
		} else {
			if counter > 0 {
				counter--
			} else {
				// insert a '('
				insertion++
				counter = 1
			}
		}
	}
	return insertion + counter
}

// @lc code=end

