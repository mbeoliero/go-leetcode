/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	less := &ListNode{Val: -1}
	l := less
	more := &ListNode{Val: -1}
	m := more
	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			m.Next = head
			m = m.Next
		}
		// temp := head.Next
		// head.Next = nil
		// head = temp
		head = head.Next
	}
	m.Next = nil
	l.Next = more.Next
	return less.Next
}

// @lc code=end

// func partition(head *ListNode, x int) *ListNode {
// 	less := &ListNode{Val: -1}
// 	l := less
// 	more := &ListNode{Val: -1}
// 	m := more
// 	for head != nil {
// 		if head.Val < x {
// 			l.Next = head
// 			l = l.Next
// 		} else {
// 			m.Next = head
// 			m = m.Next
// 		}
// 		head = head.Next
// 	}
// 	l.Next = more.Next
// 	return less.Next
// }
