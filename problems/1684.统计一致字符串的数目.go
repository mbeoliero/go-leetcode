/*
 * @lc app=leetcode.cn id=1684 lang=golang
 *
 * [1684] 统计一致字符串的数目
 */

// @lc code=start
func countConsistentStrings(allowed string, words []string) int {
	all := make(map[rune]struct{})
	for _, a := range allowed {
		all[a] = struct{}{}
	}

	count := func(word string) int {
		for _, w := range word {
			_, ok := all[w]
			if !ok {
				return 0
			}
		}
		return 1
	}

	res := 0
	for _, w := range words {
		res += count(w)
	}

	return res
}

// @lc code=end

