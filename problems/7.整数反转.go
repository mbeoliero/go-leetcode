/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
// func reverse(x int) (rev int) {
// 	for x != 0 {
// 		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
// 			return 0
// 		}
// 		digit := x % 10
// 		x /= 10
// 		rev = rev*10 + digit
// 	}
// 	return
// }

func reverse(x int) int {
	var rev func(x int, res *int)
	rev = func(x int, res *int) {
		if x == 0 {
			return
		}
		if *res < math.MinInt32/10 || *res > math.MaxInt32/10 {
			*res = 0
			return
		}

		*res = *res*10 + x%10
		rev(x/10, res)
	}

	var res int
	rev(x, &res)
	return res
}

// @lc code=end

