/*
 * @lc app=leetcode.cn id=923 lang=golang
 *
 * [923] 三数之和的多种可能
 */

// @lc code=start
func threeSumMulti(arr []int, target int) int {
	var mp map[int]int = map[int]int{0: 0}
	var res int

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			diff := target - (arr[i] + arr[j])
			res += mp[diff]
		}
		mp[arr[i]]++
	}
	return res % int(1e9+7)
}

// @lc code=end

