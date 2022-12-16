/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		idx := len(queue)
		path := []int{}

		for i := 0; i < idx; i++ {
			v := queue[i]
			if v != nil {
				if v.Left != nil {
					queue = append(queue, v.Left)
				}
				if v.Right != nil {
					queue = append(queue, v.Right)
				}
			}
			path = append(path, v.Val)
		}
		queue = queue[idx:]
		res = append(res, path)
	}
	return res
}

// @lc code=end

