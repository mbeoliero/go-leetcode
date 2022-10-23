/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	m := len(needle)
	n := len(haystack)
	dp := kmp(needle)

	j := 0
	for i := 0; i < n; i++ {
		j = dp[j][int(haystack[i])]
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

func kmp(pat string) [][]int {
	m := len(pat)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, 256)
	}

	dp[0][int(pat[0])] = 1
	x := 0

	for j := 1; j < m; j++ {
		for c := 0; c < 256; c++ {
			if int(pat[j]) == c {
				dp[j][c] = j + 1
			} else {
				dp[j][c] = dp[x][c]
			}
		}

		x = dp[x][int(pat[j])]
	}
	return dp
}

// @lc code=end

// RabinKarp
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}