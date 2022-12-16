/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	// var dfs func(root *TreeNode, sum int) (res int)
	// dfs = func(root *TreeNode, sum int) (res int) {
	// 	if root == nil {
	// 		return
	// 	}

	// 	val := root.Val
	// 	if val == sum {
	// 		res++
	// 	}

	// 	res += dfs(root.Left, sum-val)
	// 	res += dfs(root.Right, sum-val)
	// 	return
	// }

	if root == nil {
		return 0
	}
	ans := dfs(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}

func dfs(root *TreeNode, sum int) (res int) {
	if root == nil {
		return
	}

	val := root.Val
	if val == sum {
		res++
	}

	res += dfs(root.Left, sum-val)
	res += dfs(root.Right, sum-val)
	return
}

// @lc code=end

