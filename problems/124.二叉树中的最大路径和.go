/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	res := math.MinInt16
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		if left+right+root.Val > res {
			res = left + right + root.Val
		}
		if left+root.Val > right+root.Val {
			if left+root.Val < 0 {
				return 0
			}
			return left + root.Val
		} else {
			if right+root.Val < 0 {
				return 0
			}
			return right + root.Val
		}
	}
	dfs(root)
	return res
}

// @lc code=end

