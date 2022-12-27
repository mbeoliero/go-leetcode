/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	root := &TreeNode{}
	if root1 == nil {
		root.Val = root2.Val
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	} else if root2 == nil {
		root.Val = root1.Val
		root.Left = mergeTrees(root1.Left, nil)
		root.Right = mergeTrees(root1.Right, nil)
	} else {
		root.Val = root1.Val + root2.Val
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root
}

// @lc code=end

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// if one of t1 and t2 is nil, return the other
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// merge t1 and t2
	root := &TreeNode{Val: t1.Val + t2.Val}
	// recursion
	root.Left = mergeTrees(t1.Left, t2.Left)
	root.Right = mergeTrees(t1.Right, t2.Right)
	return root
}