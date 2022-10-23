package common

func KnapSack(W int, N int, weight []int, val []int) int {
	// 0-1背包问题
	// 背包承重w
	// n个物品
	// 第i个物品 重量：weight[i] 价值：val[i]

	// dp[i][w]表示前i个物品，背包容量为w时，能装下的最大价值
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-weight[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = Max(
					dp[i-1][w-weight[i-1]]+val[i-1], //放入背包
					dp[i-1][w],                      // 不放入
				)
			}
		}
	}
	return dp[N][W]
}
