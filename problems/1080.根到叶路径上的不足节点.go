/*
 * @lc app=leetcode.cn id=1080 lang=golang
 *
 * [1080] 根到叶路径上的不足节点
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
// func sufficientSubset(root *TreeNode, limit int) *TreeNode {
// 	var dfs func(root *TreeNode, sum int) *TreeNode
// 	dfs = func(root *TreeNode, sum int) *TreeNode {
// 		if root == nil {
// 			return nil
// 		}

// 		sum += root.Val
// 		root.Left = dfs(root.Left, sum)
// 		root.Right = dfs(root.Right, sum)

// 		if root.Left == nil && root.Right == nil {
// 			if sum < limit {
// 				return nil
// 			}
// 		}
// 		return root
// 	}

// 	return dfs(root, 0)
// }

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	_, root = helper(root, 0, limit)
	return root
}

func helper(root *TreeNode, sum int, limit int) (int, *TreeNode) {
	sum += root.Val
	var maxVal int
	if root.Left == nil && root.Right == nil {
		// If this is the leaf node, the maxValue is going to be the sum so far
		maxVal = sum
	} else {
		// Otherwise, the max value is going to be the max of left and right subtree values
		leftMax, rightMax := math.MinInt32, math.MinInt32
		if root.Left != nil {
			leftMax, root.Left = helper(root.Left, sum, limit)
		}
		if root.Right != nil {
			rightMax, root.Right = helper(root.Right, sum, limit)
		}
		maxVal = max(leftMax, rightMax)
	}
	// If the max value of the root to leaf path sum is less than the limit, it means
	// all path sums are less than the limit
	if maxVal < limit {
		return maxVal, nil
	}
	return maxVal, root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

