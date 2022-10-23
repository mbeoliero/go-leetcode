/*
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] 最深叶节点的最近公共祖先
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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {

	var dfs func(root *TreeNode, depth int) int
	var lca *TreeNode
	var max int

	// dfs does post-order traversal along with keeping track of
	// depths of each node and returns distance between current
	// node and its deepest leaf node
	dfs = func(root *TreeNode, depth int) int {
		// base case
		if root == nil {
			return 0
		}
		// update max depth
		if depth > max {
			max = depth
		}
		// traverse left and right children
		left := dfs(root.Left, depth+1)
		right := dfs(root.Right, depth+1)

		// refer to picture
		if left == right && left+depth == max {
			lca = root
		}
		return Max(left, right) + 1
	}
	_ = dfs(root, 0)
	return lca
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

