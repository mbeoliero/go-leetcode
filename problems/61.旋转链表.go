/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}

	k = k % n
	if k == 0 {
		return head
	}

	p = head

	for i := 0; i < k; i++ {
		p = p.Next
	}

	q := head
	for p != nil && p.Next != nil {
		p = p.Next
		q = q.Next
	}

	newHead := q.Next
	q.Next = nil
	p.Next = head
	return newHead
}

// @lc code=end

