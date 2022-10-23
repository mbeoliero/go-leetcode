/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'G':
			nums[i] = 1
		case 'C':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}

	L, R := 10, 4
	RL := int(math.Pow(float64(R), float64(L-1)))

	hash := make(map[int]bool)
	ap := make(map[int]bool)
	res := make([]string, 0)
	wHash := 0

	left, right := 0, 0
	for right < len(nums) {
		wHash = R*wHash + nums[right]
		right++

		if right-left == L {
			if hash[wHash] && !ap[wHash] {
				res = append(res, s[left:right])
				ap[wHash] = true
			} else {
				hash[wHash] = true
			}
			wHash = wHash - nums[left]*RL
			left++
		}
	}
	return res
}

// func findRepeatedDnaSequences(s string) []string {
// 	exist := make(map[string]struct{})
// 	ap := make(map[string]struct{})
// 	res := make([]string, 0)

// 	left, right := 0, 10
// 	for right <= len(s) {
// 		if _, ok := exist[s[left:right]]; ok {
// 			if _, ok := ap[s[left:right]]; !ok {
// 				res = append(res, s[left:right])
// 				ap[s[left:right]] = struct{}{}
// 			}
// 		}
// 		exist[s[left:right]] = struct{}{}
// 		left++
// 		right++
// 	}
// 	return res
// }

// @lc code=end

