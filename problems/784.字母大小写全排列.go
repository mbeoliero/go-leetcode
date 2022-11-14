/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{""}
	}
	var res []string
	c := S[len(S)-1]
	for _, str := range letterCasePermutation(S[:len(S)-1]) {
		res = append(res, str+string(c))
		if c >= 'a' && c <= 'z' {
			res = append(res, str+string(c-32))
		} else if c >= 'A' && c <= 'Z' {
			res = append(res, str+string(c+32))
		}
	}
	return res
}

// @lc code=end

func letterCasePermutation(s string) []string {
	res := make([]string, 0)
	backtrack(s, 0, []byte{}, &res)
	return res
}

func backtrack(s string, i int, path []byte, res *[]string) {
	if i == len(s) {
		*res = append(*res, string(path))
		return
	}

	if s[i] >= '0' && s[i] <= '9' {
		path = append(path, s[i])
		backtrack(s, i+1, path, res)
		path = path[:len(path)-1]
	} else {
		var other byte
		if s[i] >= 'a' && s[i] <= 'z' {
			other = s[i] - 'a' + 'A'
		} else {
			other = s[i] - 'A' + 'a'
		}

		path = append(path, s[i])
		backtrack(s, i+1, path, res)
		path = path[:len(path)-1]

		path = append(path, other)
		backtrack(s, i+1, path, res)
		path = path[:len(path)-1]
	}
}