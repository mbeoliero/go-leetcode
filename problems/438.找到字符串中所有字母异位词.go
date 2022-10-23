/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	if len(p) > len(s) {
		return res
	}
	ori, cnt := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		ori[p[i]]++
	}

	for i := 0; i < len(p); i++ {
		if _, ok := ori[s[i]]; ok {
			cnt[s[i]]++
		}
	}
	check := func() bool {
		for k, v := range ori {
			if cnt[k] != v {
				return false
			}
		}
		return true
	}
	if check() {
		res = append(res, 0)
	}

	left, right := 0, len(p)
	for right < len(s) {
		if _, ok := ori[s[right]]; ok {
			cnt[s[right]]++
		}
		if _, ok := ori[s[left]]; ok {
			cnt[s[left]]--
		}
		if check() {
			res = append(res, left+1)
		}

		left++
		right++
	}
	return res
}

// @lc code=end

