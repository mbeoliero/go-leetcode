/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	stack := make([]*Node, 0)
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		tmp := stack[:len(stack)]
		val := make([]int, 0)
		stack = stack[len(stack):]

		for _, n := range tmp {
			val = append(val, n.Val)

			for _, child := range n.Children {
				if child != nil {
					stack = append(stack, child)
				}
			}
		}

		res = append(res, val)
	}
	return res
}

// @lc code=end

