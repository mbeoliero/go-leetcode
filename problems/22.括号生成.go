/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
var res []string

func generateParenthesis(n int) []string {
	track := make([]rune, 0)
	res = make([]string, 0)

	backtrack(n, n, track)
	return res
}

func backtrack(left, right int, track []rune) {
	if left > right || left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		res = append(res, string(track))
	}

	track = append(track, '(')
	backtrack(left-1, right, track)
	track = track[:len(track)-1]

	track = append(track, ')')
	backtrack(left, right-1, track)
	track = track[:len(track)-1]
}

// @lc code=end

