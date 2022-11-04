package common

import (
	"container/heap"
	"container/list"
)

type LFUCache1 struct {
	Nodes   map[int]*list.Element
	Lists   map[int]*list.List
	Cap     int
	MinFreq int
}

type Node struct {
	Key, Val, Freq int
}

func Constructor2(capacity int) LFUCache1 {
	return LFUCache1{
		Nodes:   make(map[int]*list.Element),
		Lists:   make(map[int]*list.List),
		Cap:     capacity,
		MinFreq: 0,
	}
}

func (c *LFUCache1) Get(key int) int {
	v, ok := c.Nodes[key]
	if !ok {
		return -1
	}

	curr := v.Value.(*Node)
	c.Lists[curr.Freq].Remove(v)

	curr.Freq++
	if _, ok := c.Lists[curr.Freq]; !ok {
		c.Lists[curr.Freq] = list.New()
	}

	newList := c.Lists[curr.Freq]
	newNode := newList.PushFront(curr)
	c.Nodes[key] = newNode

	if curr.Freq-1 == c.MinFreq && c.Lists[curr.Freq-1].Len() == 0 {
		c.MinFreq++
	}
	return curr.Val
}

func (c *LFUCache1) Put(key, val int) {
	if c.Cap == 0 {
		return
	}

	if curr, ok := c.Nodes[key]; ok {
		currNode := curr.Value.(*Node)
		currNode.Val = val
		c.Get(key)
		return
	}

	if c.Cap == len(c.Nodes) {
		currList := c.Lists[c.MinFreq]
		backNode := currList.Back()
		delete(c.Nodes, backNode.Value.(*Node).Key)
		currList.Remove(backNode)
	}

	c.MinFreq = 1
	currNode := &Node{Key: key, Val: val, Freq: c.MinFreq}
	if _, ok := c.Lists[c.MinFreq]; !ok {
		c.Lists[c.MinFreq] = list.New()
	}

	newList := c.Lists[c.MinFreq]
	newNode := newList.PushFront(currNode)
	c.Nodes[key] = newNode
}

type LFUCache2 struct {
	Cap   int
	PQ    PriorityQueue
	Hash  map[int]*Item
	Count int
}

type Item struct {
	Key, Val, Freq int
	Count, Index   int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Freq == pq[j].Freq {
		return pq[i].Count < pq[j].Count
	}
	return pq[i].Freq < pq[j].Freq
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)

	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, val, freq, count int) {
	item.Val = val
	item.Count = count
	item.Freq = freq
	heap.Fix(pq, item.Index)
}

func Constructor3(capacity int) LFUCache2 {
	return LFUCache2{
		PQ:   PriorityQueue{},
		Hash: make(map[int]*Item),
		Cap:  capacity,
	}
}

func (c *LFUCache2) Get(key int) int {
	if c.Cap == 0 {
		return -1
	}

	if item, ok := c.Hash[key]; ok {
		c.Count++
		c.PQ.update(item, item.Val, item.Freq+1, c.Count)
		return item.Val
	}
	return -1
}

func (c *LFUCache2) Put(key, val int) {
	if c.Cap == 0 {
		return
	}

	if item, ok := c.Hash[key]; ok {
		c.PQ.update(item, val, item.Freq+1, c.Count)
		return
	}

	if len(c.PQ) == c.Cap {
		item := heap.Pop(&c.PQ).(*Item)
		delete(c.Hash, item.Key)
	}

	item := &Item{
		Val:   val,
		Key:   key,
		Count: c.Count,
	}

	heap.Push(&c.PQ, item)
	c.Hash[key] = item
}
