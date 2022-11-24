/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	nums := make([]int, 0)
	p := head
	for p != nil {
		nums = append(nums, p.Val)
		p = p.Next
	}

	MergeSort(nums)

	p = head
	for i := 0; i < len(nums); i++ {
		p.Val = nums[i]
		p = p.Next
	}

	return head
}

var temp []int

func MergeSort(nums []int) {
	temp = make([]int, len(nums))
	ParMergeSort(nums, 0, len(nums)-1)
}

func ParMergeSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	ParMergeSort(nums, low, mid)
	ParMergeSort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		temp[i] = nums[i]
	}

	i, j := low, mid+1
	p := low
	for i <= mid && j <= high {
		if temp[i] <= temp[j] {
			nums[p] = temp[i]
			i++
			p++
		} else {
			nums[p] = temp[j]
			j++
			p++

		}
	}

	for i <= mid {
		nums[p] = temp[i]
		i++
		p++
	}
	for j <= high {
		nums[p] = temp[j]
		j++
		p++
	}
}

// @lc code=end

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}
