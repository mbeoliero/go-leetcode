/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = fast
		}
		temp := fast
		fast = fast.Next
		temp.Next = nil
	}
	if fast != nil && slow.Val != fast.Val {
		slow.Next = fast
	}
	return head
}

// @lc code=end

