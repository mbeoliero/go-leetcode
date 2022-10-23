/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	hashMap := [26]int{}
	visited := [26]bool{}

	for i := range s {
		hashMap[int(s[i])-int('a')] = i
	}

	var res []byte
	for i := range s {
		ch := s[i]
		for len(res) > 0 {
			lastCh := res[len(res)-1]
			if ch <= lastCh && hashMap[int(lastCh)-int('a')] >= i && !visited[int(ch)-int('a')] {
				visited[int(lastCh)-int('a')] = false
				res = res[:len(res)-1]
				continue
			}
			break
		}

		if !visited[int(ch)-int('a')] {
			res = append(res, ch)
			visited[int(ch)-int('a')] = true
		}
	}
	return string(res)
}

// func removeDuplicateLetters(s string) string {
// 	stack := make([]rune, 0)
// 	inStack := [256]bool{}

// 	count := [256]int{}
// 	for _, r := range s {
// 		count[r]++
// 	}

// 	for _, r := range s {
// 		count[r]--
// 		if inStack[r] {
// 			continue
// 		}

// 		for len(stack) > 0 && r < stack[len(stack)-1] && count[stack[len(stack)-1]] > 0 {
// 			inStack[stack[len(stack)-1]] = false
// 			stack = stack[:len(stack)-1]
// 		}
// 		stack = append(stack, r)
// 		inStack[r] = true
// 	}
// 	return string(stack)
// }

// @lc code=end

