/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		length := len(stack)
		for idx, s := range stack {
			if idx == length-1 {
				s.Next = nil
			} else {
				s.Next = stack[idx+1]
			}

			if s.Left != nil && s.Right != nil {
				stack = append(stack, s.Left, s.Right)
			}
		}

		stack = stack[length:]
	}

	return root
}

// @lc code=end

