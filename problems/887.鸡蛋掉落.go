/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 扔m次，m小于n
	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}

	return m
}

// @lc code=end

var memo [][]int

func superEggDrop(k int, n int) int {
	memo = make([][]int, k+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dp(k, n)
}

// k个鸡蛋，n层楼
func dp(k, n int) int {
	if k == 1 {
		return n
	}
	if n == 0 {
		return 0
	}

	if memo[k][n] != -1 {
		return memo[k][n]
	}

	res := math.MaxInt16
	// for i := 1; i < n+1; i++ {
	// 	// res = min(res,
	// 	// 	max( //最坏情况
	// 	// 		dp(k, n-i),   //鸡蛋没碎
	// 	// 		dp(k-1, i-1), //鸡蛋碎了
	// 	// 	)+1, // 扔一次鸡蛋
	// 	// )

	// 	// 求两个函数的相交点 (x型)
	//  // dp(k,n-i): 递减函数
	//  // dp(k-1,i-1): 递增函数
	// 	if dp(k, n-i) == dp(k-1, i-1) {
	// 		return dp(k, n-i)
	// 	}
	// }
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2
		broken := dp(k-1, mid-1)  // 碎
		notBroken := dp(k, n-mid) // 不碎

		if broken > notBroken {
			high = mid - 1
			res = min(res, broken+1)
		} else {
			low = mid + 1
			res = min(res, notBroken+1)
		}
	}

	memo[k][n] = res
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
