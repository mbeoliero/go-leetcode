/*
 * @lc app=leetcode.cn id=817 lang=golang
 *
 * [817] 链表组件
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
	mp := make(map[int]bool)
	for _, n := range nums {
		mp[n] = true
	}

	count := 0
	for head != nil {
		if mp[head.Val] {
			count++
			for head != nil && mp[head.Val] {
				head = head.Next
			}
		}

		if head != nil {
			head = head.Next
		}
	}

	return count
}

// @lc code=end

