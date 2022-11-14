/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	var res int
	var dfs func(root *TreeNode, num int)

	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}

		num = num*10 + root.Val
		if root.Left == nil && root.Right == nil {
			res += num
			return
		}

		dfs(root.Left, num)
		dfs(root.Right, num)
	}

	dfs(root, 0)
	return res
}

// @lc code=end

