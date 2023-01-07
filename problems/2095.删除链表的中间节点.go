/*
 * @lc app=leetcode.cn id=2095 lang=golang
 *
 * [2095] 删除链表的中间节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {return nil}
    
    pre, slow, fast := head, head, head
    
    for slow.Next != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next
        
        if fast.Next != nil {
            fast = fast.Next
        }
    }
    
    pre.Next = slow.Next
    
    return head
}
// @lc code=end

