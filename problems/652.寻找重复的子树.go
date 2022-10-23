/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)

	record := make(map[string][]*TreeNode)

	serialize(root, record)
	for _, val := range record {
		if len(val) > 1 {
			res = append(res, val[0])
		}
	}

	return res
}

func serialize(root *TreeNode, record map[string][]*TreeNode) string {
	if root == nil {
		return ""
	}

	left := serialize(root.Left, record)
	right := serialize(root.Right, record)

	key := "[" + left + "," + fmt.Sprintf("%v", root.Val) + "," + right + "]"
	record[key] = append(record[key], root)

	return key
}

// @lc code=end

