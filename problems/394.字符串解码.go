/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	var dfs func(start int) (string, int)
	dfs = func(start int) (string, int) {
		ans := ""
		dup := 0
		i := start
		for i < len(s) {
			ch := s[i]
			if '0' <= ch && ch <= '9' {
				dup = dup*10 + int(ch-'0')
			} else if ch == '[' {
				tmp, end := dfs(i + 1)
				ans += strings.Repeat(tmp, dup)
				dup = 0
				i = end
			} else if 'a' <= ch && ch <= 'z' {
				ans += string(ch)
			} else if ch == ']' {
				return ans, i
			}
			i++
		}
		return ans, i
	}

	r, _ := dfs(0)
	return r
}

// @lc code=end

