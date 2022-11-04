/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dump := &ListNode{Next: head}
	prev := dump

	for head != nil {
		cnt := k - 1
		h := head

		for cnt > 0 && head != nil {
			head = head.Next
			cnt--
		}

		t := head
		if cnt == 0 && head != nil {
			next := t.Next
			he, ta := Reverse(h, t)

			prev.Next = he
			ta.Next = next
			head = next
			prev = ta
		}
	}

	return dump.Next
}

func Reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev

		prev = p
		p = next
	}
	return tail, head
}

// @lc code=end

func reverseKGroup(head *ListNode, k int) *ListNode {
	node, cnt := head, 0
	for cnt < k {
		if node == nil {
			return head
		}
		node = node.Next
		cnt++
	}

	prev := reverseKGroup(node, k)
	for cnt > 0 {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
		cnt--
	}

	return prev
}
