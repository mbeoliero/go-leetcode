/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start

var memo map[string]bool

func isMatch(s string, p string) bool {
	i, j := 0, 0
	memo = make(map[string]bool)
	return dp(s, i, p, j)
}

func dp(s string, i int, p string, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	if i == len(s) { // p=x*y*z*
		if (len(p)-j)%2 != 0 {
			return false
		}

		for ; j+1 < len(p); j += 2 {
			if p[j+1] != '*' {
				return false
			}
		}
		return true
	}

	key := string(i) + "," + string(j)
	if v, ok := memo[key]; ok {
		return v
	}

	var res bool
	if s[i] == p[j] || p[j] == '.' {
		if j < len(p)-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2) || // 匹配一次
				dp(s, i+1, p, j) // 匹配多次
		} else {
			res = dp(s, i+1, p, j+1)
		}
	} else {
		if j < len(p)-1 && p[j+1] == '*' {
			res = dp(s, i, p, j+2) // 匹配零次
		} else {
			res = false
		}
	}

	memo[key] = res
	return res
}

// @lc code=end

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	var dp [][]bool
	var t []bool
	for i := 0; i <= slen; i++ {
		t = make([]bool, plen+1)
		dp = append(dp, t)
	}

	// dp[i][j] holds a flag that whether s[0:i] mathes to p[0:j].
	// (Note that we take the 0th index as "empty" string case)
	for i := 0; i <= slen; i++ {
		for j := 0; j <= plen; j++ {
			if i == 0 && j == 0 {
				// in case both s and p are empty
				dp[i][j] = true
				continue
			} else if i == 0 {
				// in case only s is empty.
				// To match an empty string, p should be like a*, a*b*, a*.*a*, ....
				// (* should appear in all odd indexes of p)
				dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
				continue
			} else if j == 0 {
				// in case only p is mepty.
				// No string s cannot match to empty regular expression
				// unless s itself is empty too (which is already handled)
				dp[i][j] = false
				continue
			}

			// in case both s and p are non-empty
			// consider what should be sufficed to be able to say s[i-1] matches to p[j-1]
			// (= dp[i][j] is true).
			if p[j-1] != '*' {
				// if p[j-1] is not '*', then
				// 1) p[j-1] and s[j-1] should match (same character or p[j-1] == '.')
				// 2) p[0:j-1] should match to s[0:i-1]
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				// if p[j-1] is '*', we need to consider two scenarios
				// case1) consume the '*'
				// to consume * and have s[:i] match to p[:j],
				// p[j-2] should equal to s[i-1] (or p[j-2] can be '.' to be equal to any s[i-1])
				// AND s[:i-1] should match to p[:j] as well.
				//                 ↓ i-1
				// s: [ ] [ ] [a] [a]
				// p: [ ] [a] [*] [ ]
				//             ↑ j-1
				//
				//                 ↓ i-1
				// s: [ ] [ ] [b] [a]
				// p: [ ] [.] [*] [ ]
				//             ↑ j-1
				// in other words, because '*' is used with its previous character,
				// for s[i-1] to be included as "a*" or ".*" in p,
				// s[i-2] should also be included in "a*" or ".*".
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i-1][j]
				}
				// case2) skip the '*'
				//                 ↓ i-1
				// s: [ ] [ ] [b] [a]
				// p: [a] [b] [*] [ ]
				//     ↑ j-3   ↑ j-1
				// even if we can't fulfill the condition above, if s[:i-1] matches to p[j-3],
				// we can say s[:i-1] also matches to p[j-1]
				// because '*' can mean "zero" appearance of the previous character.
				// in other words, we can just skip "b*", not using it.
				if dp[i][j-2] {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[slen][plen]
}
