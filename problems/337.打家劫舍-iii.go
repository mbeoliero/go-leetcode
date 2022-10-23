/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
	res := dp(root)
	return max(res[0], res[1])
}

func dp(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dp(root.Left)
	right := dp(root.Right)

	do := root.Val + left[0] + right[0]
	not := max(left[0], left[1]) + max(right[0], right[1])
	return []int{not, do}
}

// func rob(root *TreeNode) int {
// 	memo := make(map[*TreeNode]int)
// 	var dp func(root *TreeNode) int
// 	dp = func(root *TreeNode) int {
// 		if root == nil {
// 			return 0
// 		}
// 		if v, ok := memo[root]; ok {
// 			return v
// 		}

// 		do := root.Val
// 		if root.Left != nil {
// 			do += dp(root.Left.Left) + dp(root.Left.Right)
// 		}
// 		if root.Right != nil {
// 			do += dp(root.Right.Left) + dp(root.Right.Right)
// 		}

// 		not := dp(root.Left) + dp(root.Right)
// 		res := max(not, do)
// 		memo[root] = res
// 		return res
// 	}
// 	return dp(root)
// }

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

