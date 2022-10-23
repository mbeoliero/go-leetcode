/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	if len(preorder) == 0 || len(inorder) == 0 {
// 		return nil
// 	}

// 	val := preorder[0]
// 	preIdx := make(map[int]int)
// 	preidx, inidx := len(preorder), len(preorder)

// 	for i, n := range preorder {
// 		preIdx[n] = i
// 	}

// 	for i, n := range inorder {
// 		if n == val {
// 			inidx = i
// 		}
// 		if i > inidx {
// 			if preidx > preIdx[n] {
// 				preidx = preIdx[n]
// 			}
// 		}
// 	}

// 	root := &TreeNode{
// 		Val:   val,
// 		Left:  buildTree(preorder[1:preidx], inorder[:inidx]),
// 		Right: buildTree(preorder[preidx:], inorder[inidx+1:]),
// 	}
// 	return root
// }
func buildTree(preorder []int, inorder []int) *TreeNode {
	inIdx := make(map[int]int)
	for i, n := range inorder {
		inIdx[n] = i
	}

	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inIdx)
}

func build(preorder []int, preLeft, preRight int, inorder []int, inLeft, inRight int, idx map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}

	val := preorder[preLeft]
	index := idx[val]
	leftSize := index - inLeft

	root := &TreeNode{
		Val:   val,
		Left:  build(preorder, preLeft+1, preLeft+leftSize, inorder, inLeft, index-1, idx),
		Right: build(preorder, preLeft+leftSize+1, preRight, inorder, index+1, inRight, idx),
	}
	return root
}

// @lc code=end

