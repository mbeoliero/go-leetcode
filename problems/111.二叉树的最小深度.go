/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			head := queue[i]
			if head.Left == nil && head.Right == nil {
				return depth
			}
			if head.Left != nil {
				queue = append(queue, head.Left)
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		queue = queue[size:]
		depth++
	}
	return depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

