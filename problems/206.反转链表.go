/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	p, q := head, head.Next
// 	p.Next = nil
// 	var temp *ListNode
// 	for q != nil {
// 		temp = q.Next
// 		q.Next = p
// 		p = q
// 		q = temp
// 	}
// 	return p
// }

// @lc code=end

