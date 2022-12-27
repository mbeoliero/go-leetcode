/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type LRUCache struct {
	KVMap map[int]int
	KeyEl map[int]*list.Element
	Cap   int
	List  *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		KVMap: make(map[int]int),
		Cap:   capacity,
		List:  list.New(),
		KeyEl: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.KVMap[key]; !ok {
		return -1
	}

	element := this.KeyEl[key]
	this.List.MoveToFront(element)
	return this.KVMap[key]
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.KVMap[key]; ok {
		this.KVMap[key] = value
		element := this.KeyEl[key]
		this.List.MoveToFront(element)
		return
	}

	if this.List.Len() >= this.Cap {
		e := this.List.Back()
		key := e.Value.(int)
		this.List.Remove(this.List.Back())
		delete(this.KVMap, key)
		delete(this.KeyEl, key)
	}

	this.KVMap[key] = value
	this.List.PushFront(key)
	this.KeyEl[key] = this.List.Front()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

