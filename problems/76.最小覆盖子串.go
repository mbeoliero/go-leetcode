/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

// func minWindow(s string, t string) string {
// 	ori, cnt := make(map[byte]int), make(map[byte]int)
// 	for i := 0; i < len(t); i++ {
// 		ori[t[i]]++
// 	}

// 	check := func() bool {
// 		for k, v := range ori {
// 			if cnt[k] < v {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	if len(t) > len(s) || len(s) == 0 {
// 		return ""
// 	}

// 	res := s + "]"
// 	left, right := 0, 0
// 	for right < len(s) {
// 		if right < len(s) && ori[s[right]] > 0 {
// 			cnt[s[right]]++
// 		}
// 		for check() && left <= right {
// 			if right+1 <= len(s) && len(res) > right-left {
// 				res = s[left : right+1]
// 			}
// 			if _, ok := ori[s[left]]; ok {
// 				cnt[s[left]] -= 1
// 			}
// 			left++
// 		}
// 		right++
// 	}
// 	if len(res) > len(s) {
// 		return ""
// 	}
// 	return res
// }

// @lc code=end

