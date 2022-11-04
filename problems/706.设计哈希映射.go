/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
type Node struct {
	Key  int
	Val  int
	Next *Node
}

type MyHashMap struct {
	Items []*Node
	Size  int
}

func Constructor() MyHashMap {
	size := 1000
	return MyHashMap{
		Items: make([]*Node, size),
		Size:  size,
	}
}

func (this *MyHashMap) Put(key int, value int) {
	index := key % this.Size
	if this.Items[index] == nil {
		this.Items[index] = &Node{Key: key, Val: value}
		return
	}

	item := this.Items[index]
	for item != nil {
		if item.Key == key {
			item.Val = value
			return
		}

		if item.Next == nil {
			item.Next = &Node{Key: key, Val: value}
			return
		}

		item = item.Next
	}
}

func (this *MyHashMap) Get(key int) int {
	index := key % this.Size
	if this.Items[index] == nil {
		return -1
	}

	item := this.Items[index]
	for item != nil {
		if item.Key == key {
			return item.Val
		}

		item = item.Next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := key % this.Size
	if this.Items[index] == nil {
		return
	}

	item := this.Items[index]
	if item.Key == key {
		this.Items[index] = item.Next
		return
	}

	prev, item := item, item.Next
	for item != nil {
		if item.Key == key {
			prev.Next = item.Next
			return
		}

		prev, item = item, item.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

