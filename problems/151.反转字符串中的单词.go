/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	strList := strings.Fields(s)
	reverseString(strList, 0, len(strList)-1)
	return strings.Join(strList, " ")
}

func reverseString(s []string, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// func reverseWords(s string) string {
// 	r := []rune(s)
// 	reverseString(r, 0, len(s)-1)
// 	left, right := 0, 0

// 	r = append(r, 32)
// 	for right < len(r) {
// 		if r[right] == 32 {
// 			reverseString(r, left, right-1)
// 			left = right
// 		}
// 		right++
// 		for left < right {
// 			if r[left] == 32 {
// 				left++
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return string(r[:len(r)-1])
// }

// func reverseString(s []rune, left, right int) {
// 	for left < right {
// 		s[left], s[right] = s[right], s[left]
// 		left++
// 		right--
// 	}
// }

// func reverseString(s string) string {
// 	res := make([]rune, 0, len(s))
// 	func() {
// 		for _, v := range s {
// 			defer func(r rune) {
// 				res = append(res, r)
// 			}(v)
// 		}
// 	}()
// 	return string(res)
// }

// @lc code=end

