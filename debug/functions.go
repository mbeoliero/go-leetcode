package debug

import "leetcode/common"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 1
	idx := make(map[byte]int)
	left, right := 0, 0
	for ; right < len(s); right++ {
		tag := s[right]
		if next, ok := idx[tag]; ok {
			if next+1 > left {
				left = next + 1
			}
		}
		idx[tag] = right
		res = common.Max(res, right-left+1)
	}
	return res
}
