/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */

// @lc code=start
func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func removeInvalidParentheses(s string) (ans []string) {
	curSet := map[string]struct{}{s: {}}
	for {
		for str := range curSet {
			if isValid(str) {
				ans = append(ans, str)
			}
		}
		if len(ans) > 0 {
			return
		}
		nextSet := map[string]struct{}{}
		for str := range curSet {
			for i, ch := range str {
				if i > 0 && byte(ch) == str[i-1] {
					continue
				}
				if ch == '(' || ch == ')' {
					nextSet[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		curSet = nextSet
	}
}

// @lc code=end

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func helper(ans *[]string, str string, start, lremove, rremove int) {
	if lremove == 0 && rremove == 0 {
		if isValid(str) {
			*ans = append(*ans, str)
		}
		return
	}

	for i := start; i < len(str); i++ {
		if i != start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if lremove+rremove > len(str)-i {
			return
		}
		// 尝试去掉一个左括号
		if lremove > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, lremove-1, rremove)
		}
		// 尝试去掉一个右括号
		if rremove > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, lremove, rremove-1)
		}
	}
}

func removeInvalidParentheses(s string) (ans []string) {
	lremove, rremove := 0, 0
	for _, ch := range s {
		if ch == '(' {
			lremove++
		} else if ch == ')' {
			if lremove == 0 {
				rremove++
			} else {
				lremove--
			}
		}
	}

	helper(&ans, s, 0, lremove, rremove)
	return
}
