package problems

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res, _ := addTwoNumbersSub(l1, l2, 0)
	return res
}

func addTwoNumbersSub(l1 *ListNode, l2 *ListNode, sub int) (*ListNode, int) {
	if l1 == nil && l2 == nil {
		if sub == 1 {
			return &ListNode{Val: 1, Next: nil}, 0
		}
		return nil, 0
	}
	if l1 == nil {
		next, _ := addTwoNumbersSub(nil, l2.Next, (l2.Val+sub)/10)
		return &ListNode{Val: (l2.Val + sub) % 10, Next: next}, (l2.Val + sub) / 10
	}
	if l2 == nil {
		next, _ := addTwoNumbersSub(nil, l1.Next, (l1.Val+sub)/10)
		return &ListNode{Val: (l1.Val + sub) % 10, Next: next}, (l1.Val + sub) / 10
	}

	next, _ := addTwoNumbersSub(l1.Next, l2.Next, (l1.Val+l2.Val+sub)/10)
	return &ListNode{Val: (l1.Val + l2.Val + sub) % 10, Next: next}, (l1.Val + l2.Val + sub) / 10
}

// @lc code=end
