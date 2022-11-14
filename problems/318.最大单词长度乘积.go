/*
 * @lc app=leetcode.cn id=318 lang=golang
 *
 * [318] 最大单词长度乘积
 */

// @lc code=start
func maxProduct(words []string) int {
	mask := make([]int, len(words))

	for i, word := range words {
		for _, w := range word {
			mask[i] |= 1 << (w - 'a')
		}
	}

	ans := 0
	for i := 0; i < len(mask); i++ {
		for j := i + 1; j < len(mask); j++ {
			if mask[i]&mask[j] == 0 {
				if ans < len(words[i])*len(words[j]) {
					ans = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return ans
}

// @lc code=end

