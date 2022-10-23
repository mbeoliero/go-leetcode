/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	l, count := 0, [26]int{}
	for i := range s1 {
		count[s1[i]-97]++
	}

	for r := range s2 {
		count[s2[r]-97]--
		if count == [26]int{} {
			return true
		}

		if r+1 >= len(s1) {
			count[s2[l]-97]++
			l++
		}
	}
	return false
}

// func checkInclusion(s1 string, s2 string) bool {
// 	if len(s1) > len(s2) {
// 		return false
// 	}
// 	ori, cnt := make(map[byte]int), make(map[byte]int)
// 	for i := 0; i < len(s1); i++ {
// 		ori[s1[i]]++
// 	}

// 	for i := 0; i < len(s1); i++ {
// 		if _, ok := ori[s2[i]]; ok {
// 			cnt[s2[i]]++
// 		}
// 	}
// 	check := func() bool {
// 		for k, v := range ori {
// 			if cnt[k] != v {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	if check() {
// 		return true
// 	}

// 	left, right := 0, len(s1)
// 	for right < len(s2) {
// 		if _, ok := ori[s2[right]]; ok {
// 			cnt[s2[right]]++
// 		}
// 		if _, ok := ori[s2[left]]; ok {
// 			cnt[s2[left]]--
// 		}
// 		if check() {
// 			return true
// 		}

// 		left++
// 		right++
// 	}
// 	return false
// }

// @lc code=end

