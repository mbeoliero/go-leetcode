/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func Mins(value int, values ...int) int {
	for _, v := range values {
		if v < value {
			value = v
		}
	}
	return value
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			offset := 1
			if word1[i-1] == word2[j-1] {
				offset = 0
			}
			dp[i][j] = Mins(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+offset)
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

// var memo [][]int

// func minDistance(word1 string, word2 string) int {
// 	len1, len2 := len(word1), len(word2)
// 	memo = make([][]int, len1)
// 	for i := 0; i < len1; i++ {
// 		arr := make([]int, len2)
// 		for j := 0; j < len2; j++ {
// 			arr[j] = -1
// 		}
// 		memo[i] = arr
// 	}

// 	return dp(word1, len1-1, word2, len2-1)
// }

// func dp(word1 string, idx1 int, word2 string, idx2 int) int {
// 	if idx1 == -1 {
// 		return idx2 + 1
// 	}
// 	if idx2 == -1 {
// 		return idx1 + 1
// 	}
// 	if memo[idx1][idx2] != -1 {
// 		return memo[idx1][idx2]
// 	}

// 	if word1[idx1] == word2[idx2] {
// 		memo[idx1][idx2] = dp(word1, idx1-1, word2, idx2-1)
// 	} else {
// 		ins := dp(word1, idx1, word2, idx2-1) + 1
// 		del := dp(word1, idx1-1, word2, idx2) + 1
// 		swap := dp(word1, idx1-1, word2, idx2-1) + 1
// 		memo[idx1][idx2] = min(ins, del, swap)
// 	}
// 	return memo[idx1][idx2]
// }

// func min(val ...int) int {
// 	m := math.MaxInt16
// 	for _, v := range val {
// 		if v < m {
// 			m = v
// 		}
// 	}
// 	return m
// }

// @lc code=end

var memo [][]int

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	memo = make([][]int, len1)
	for i := 0; i < len1; i++ {
		arr := make([]int, len2)
		for j := 0; j < len2; j++ {
			arr[j] = -1
		}
		memo[i] = arr
	}

	return dp(word1, len1-1, word2, len2-1)
}

func dp(word1 string, idx1 int, word2 string, idx2 int) int {
	if idx1 == -1 {
		return idx2 + 1
	}
	if idx2 == -1 {
		return idx1 + 1
	}
	if memo[idx1][idx2] != -1 {
		return memo[idx1][idx2]
	}

	if word1[idx1] == word2[idx2] {
		memo[idx1][idx2] = dp(word1, idx1-1, word2, idx2-1)
	} else {
		ins := dp(word1, idx1, word2, idx2-1) + 1
		del := dp(word1, idx1-1, word2, idx2) + 1
		swap := dp(word1, idx1-1, word2, idx2-1) + 1
		memo[idx1][idx2] = min(ins, del, swap)
	}
	return memo[idx1][idx2]
}

func min(val ...int) int {
	m := math.MaxInt16
	for _, v := range val {
		if v < m {
			m = v
		}
	}
	return m
}
