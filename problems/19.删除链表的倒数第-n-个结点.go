/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 1 && head.Next == nil {
		return nil
	}

	dump := &ListNode{Next: head}
	left, right := dump, dump
	for i := 1; i <= n; i++ {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	tmp := left.Next
	left.Next = tmp.Next
	tmp.Next = nil
	return dump.Next
}

// @lc code=end