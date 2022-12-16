/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var phone = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := make([]string, 0)
	var dfs func(i int, path []rune)
	dfs = func(i int, path []rune) {
		if i == len(digits) {
			res = append(res, string(path))
			return
		}

		str := phone[digits[i]]
		for _, v := range str {
			path = append(path, v)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []rune{})
	return res
}

// @lc code=end

