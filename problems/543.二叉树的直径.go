/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 计算每个结点左右孩子的高度
		left := dfs(node.Left)
		right := dfs(node.Right)

		// 路径长度等于左右孩子高度之和
		if left+right > res {
			res = left + right
		}

		// 返回较高的结点作为本结点的高度
		if left > right {
			return left + 1
		}
		return right + 1
	}

	dfs(root)
	return res
}

// func diameterOfBinaryTree(root *TreeNode) int {
// 	return maxDepth(root)
// }

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// @lc code=end

