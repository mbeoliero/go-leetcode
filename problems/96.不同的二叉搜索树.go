/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
var memo [][]int

func numTrees(n int) int {
	// 计算闭区间 [1, n] 组成的 BST 个数
	memo = make([][]int, 0)
	return count(1, n)
}

/* 计算闭区间 [lo, hi] 组成的 BST 个数 */
func count(lo, hi int) int {
	// base case
	if lo > hi {
		return 1
	}
	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}

	var res = 0
	for i := lo; i <= hi; i++ {
		// i 的值作为根节点 root
		left := count(lo, i-1)
		right := count(i+1, hi)
		// 左右子树的组合数乘积是 BST 的总数
		res += left * right
	}
	memo[lo][hi] = res

	return res
}

// @lc code=end

