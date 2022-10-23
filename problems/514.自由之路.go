/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */

// @lc code=start

var memo [][]int
var keyMap map[byte][]int

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	keyMap = make(map[byte][]int)
	for i := 0; i < m; i++ {
		if _, ok := keyMap[ring[i]]; !ok {
			keyMap[ring[i]] = make([]int, 0)
		}
		keyMap[ring[i]] = append(keyMap[ring[i]], i)
	}

	return dp(ring, 0, key, 0)
}

func dp(ring string, i int, key string, j int) int {
	if j == len(key) {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	n := len(ring)
	res := math.MaxInt16

	for _, k := range keyMap[key[j]] {
		// 转动次数
		delta := abs(k - i)
		delta = min(delta, n-delta)

		sub := dp(ring, k, key, j+1)
		res = min(res, 1+delta+sub)
	}

	memo[i][j] = res
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func findRotateSteps(ring string, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return min(dp[m-1]...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
