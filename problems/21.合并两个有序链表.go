/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	res := &ListNode{Val: -1}
	dump := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next
	}

	if list1 == nil {
		res.Next = list2
	}
	if list2 == nil {
		res.Next = list1
	}
	return dump.Next
}

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil && list2 == nil {
// 		return nil
// 	}
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val < list2.Val {
// 		return &ListNode{
// 			Val:  list1.Val,
// 			Next: mergeTwoLists(list1.Next, list2),
// 		}
// 	}
// 	return &ListNode{
// 		Val:  list2.Val,
// 		Next: mergeTwoLists(list1, list2.Next),
// 	}
// }

// @lc code=end

