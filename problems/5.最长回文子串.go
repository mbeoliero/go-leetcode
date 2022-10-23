/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		if len(findPalindrome(s, i, i)) > len(res) {
			res = findPalindrome(s, i, i)
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			if len(findPalindrome(s, i, i+1)) > len(res) {
				res = findPalindrome(s, i, i+1)
			}
		}
	}
	return res
}

func findPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return s[l+1 : r]
		}
		l--
		r++
	}
	return s[l+1 : r]
}

// @lc code=end

