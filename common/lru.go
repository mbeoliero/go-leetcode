package common

type LinkNode struct {
	Next, Prev *LinkNode
	Key, Val   int
}

type DoubleLink struct {
	Head, Tail *LinkNode
	Size       int
}

func NewDoubleLink() *DoubleLink {
	head, tail := new(LinkNode), new(LinkNode)
	head.Next, tail.Prev = tail, head
	return &DoubleLink{
		Head: head,
		Tail: tail,
		Size: 0,
	}
}

func (d *DoubleLink) AddLast(node *LinkNode) {
	node.Prev = d.Tail.Prev
	node.Next = d.Tail
	d.Tail.Prev.Next = node
	d.Tail.Prev = node
	d.Size++
}

func (d *DoubleLink) Remove(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	d.Size--
}

func (d *DoubleLink) RemoveFirst() *LinkNode {
	if d.Head.Next == d.Tail {
		return nil
	}

	first := d.Head.Next
	d.Remove(first)
	return first
}

func (d *DoubleLink) GetSize() int {
	return d.Size
}

type LRUCache struct {
	HashMap map[int]*LinkNode
	Cache   *DoubleLink
	Cap     int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Cap:     capacity,
		HashMap: make(map[int]*LinkNode),
		Cache:   NewDoubleLink(),
	}
}

func (l *LRUCache) MakeRecent(key int) {
	node := l.HashMap[key]
	l.Cache.Remove(node)
	l.Cache.AddLast(node)
}

func (l *LRUCache) AddRecent(key, val int) {
	node := new(LinkNode)
	node.Key, node.Val = key, val
	l.Cache.AddLast(node)
	l.HashMap[key] = node
}

func (l *LRUCache) DeleteKey(key int) {
	node := l.HashMap[key]
	l.Cache.Remove(node)
	delete(l.HashMap, key)
}

func (l *LRUCache) RemoveLastRecent() {
	node := l.Cache.RemoveFirst()
	key := node.Key
	delete(l.HashMap, key)
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.HashMap[key]; !ok {
		return -1
	}

	l.MakeRecent(key)
	return l.HashMap[key].Val
}

func (l *LRUCache) Put(key, val int) {
	if _, ok := l.HashMap[key]; ok {
		delete(l.HashMap, key)
		l.AddRecent(key, val)
		return
	}

	if l.Cap == l.Cache.GetSize() {
		l.RemoveLastRecent()
	}

	l.AddRecent(key, val)
}
