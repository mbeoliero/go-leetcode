/*
 * @lc app=leetcode.cn id=791 lang=golang
 *
 * [791] 自定义字符串排序
 */

// @lc code=start
func customSortString(order string, s string) string {
	mp := make(map[rune]int)
	for i, o := range order {
		mp[o] = i
	}

	tmp := []rune(s)
	sort.Slice(tmp, func(i, j int) bool {
		if _, ok := mp[tmp[i]]; !ok {
			return false
		}

		if _, ok := mp[tmp[j]]; !ok {
			return true
		}

		return mp[tmp[i]] < mp[tmp[j]]
	})
	return string(tmp)
}

// @lc code=end

