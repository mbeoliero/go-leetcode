package problems

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodes []*ListNode

func (l ListNodes) Len() int {
	return len(l)
}

func (l ListNodes) Less(i, j int) bool {
	return l[i].Val <= l[j].Val
}

func (l ListNodes) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodes) Pop() interface{} {
	n := len(*l)
	x := (*l)[n-1]
	*l = (*l)[:n-1]
	return x
}

func (l *ListNodes) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}

var _ heap.Interface = &ListNodes{}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	pre := dummy
	hp := &ListNodes{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(hp, lists[i])
		}
	}
	for len(*hp) > 0 {
		top := heap.Pop(hp).(*ListNode)
		pre.Next = top
		pre = pre.Next
		if top.Next != nil {
			heap.Push(hp, top.Next)
		}
	}
	return dummy.Next
}

// func mergeKLists(lists []*ListNode) *ListNode {
// 	// return mergeLists(lists, 0, len(lists)-1)
// 	length := len(lists)
// 	if length == 0 {
// 		return nil
// 	}
// 	if length == 1 {
// 		return lists[0]
// 	}
// 	mid := length / 2
// 	left := mergeKLists(lists[:mid])
// 	right := mergeKLists(lists[mid:])
// 	return mergeTwoLists(left, right)
// }

// func mergeLists(lists []*ListNode, left, right int) *ListNode {
// 	if left > right {
// 		return nil
// 	}
// 	if left+1 == right {
// 		return mergeTwoLists(lists[left], lists[right])
// 	}
// 	if left == right {
// 		return lists[left]
// 	}
// 	mid := (left + right) / 2
// 	return mergeTwoLists(
// 		mergeLists(lists, left, mid),
// 		mergeLists(lists, mid+1, right),
// 	)
// }

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

// @lc code=end
