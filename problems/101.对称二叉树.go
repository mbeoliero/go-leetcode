/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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

// func isSymmetric(root *TreeNode) bool {
// 	t := traverse(root)
// 	return root == t
// }

// func traverse(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return root
// 	}

// 	traverse(root.Left)
// 	traverse(root.Right)
// 	root.Left, root.Right = root.Right, root.Left
// 	return root
// }

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Right, node2.Left) && isMirror(node1.Left, node2.Right)
}

// @lc code=end

