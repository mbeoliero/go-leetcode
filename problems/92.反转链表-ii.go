/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	head, dummy.Next = dummy, head
	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	newHead, _ := reverseList(head.Next, n-m+1)
	head.Next = newHead
	// if m == 1 { return head.Next } else { return dummy.Next }
	return dummy.Next
}

// return new head and the head of the rest
func reverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
	if cnt == 1 {
		return head, head.Next
	}
	newHead, restHead := reverseList(head.Next, cnt-1)
	head.Next.Next = head
	head.Next = restHead
	return newHead, restHead
}

// @lc code=end

