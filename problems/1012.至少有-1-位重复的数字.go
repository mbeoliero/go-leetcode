/*
 * @lc app=leetcode.cn id=1012 lang=golang
 *
 * [1012] 至少有 1 位重复的数字
 */

// @lc code=start
func numDupDigitsAtMostN(n int) int {
	if n < 10 {
		return 0
	}

	sum := 9
	i := 2
	res := 100
	for res < n {
		sum += getNotDup(i)
		res = int(math.Pow(10, float64(i)))
		i++
	}

	return n - sum
}

func getNotDupWithN(n, i int) int {
	top := n / int(math.Pow(10, float64(i-1)))
	res := 0
	temp := top - 1
	j := 9
	k := i

	for k > 1 {
		temp *= j
		j--
		k--
	}

	res += temp
}

func getNotDup(i int) int {
	res := 9
	j := 9

	for i > 1 {
		res *= j
		j--
		i--
	}

	return res
}

// @lc code=end

func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)

	dp := make([][1 << 10]int, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(int, int, bool, bool) int
	f = func(i, mask int, isLimit, isNum bool) (res int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}

		if !isLimit && isNum {
			dv := &dp[i][mask]
			if *dv >= 0 {
				return *dv
			}

			defer func() { *dv = res }()
		}

		if !isNum {
			res += f(i+1, mask, false, false)
		}

		d := 1
		if isNum {
			d = 0
		}

		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}

		for ; d <= up; d++ {
			if mask>>d&1 == 0 {
				res += f(i+1, mask|1<<d, isLimit && d == up, true)
			}
		}
		return
	}

	return n - f(0, 0, true, false)
}

