/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return tranverse(head)
}

func tranverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := tranverse(right.Next)
	res = res && left.Val == right.Val
	left = left.Next
	return res
}

// @lc code=end

