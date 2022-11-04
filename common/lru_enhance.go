package common

import (
	"container/list"
	"crypto/sha1"
	"fmt"
	"sync"
)

type LRUCache1 struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

type Pair1 struct {
	Key, Val int
}

func Constructor1(capacity int) LRUCache1 {
	return LRUCache1{
		Cap:  capacity,
		Keys: make(map[int]*list.Element),
		List: list.New(),
	}
}

func (c *LRUCache1) Get(key int) int {
	if v, ok := c.Keys[key]; ok {
		c.List.MoveToFront(v)
		return v.Value.(Pair1).Val
	}
	return -1
}

func (c *LRUCache1) Put(key, val int) {
	if v, ok := c.Keys[key]; ok {
		v.Value = Pair1{Key: key, Val: val}
		c.List.MoveToFront(v)
	} else {
		v := c.List.PushFront(Pair1{Key: key, Val: val})
		c.Keys[key] = v
	}

	if c.List.Len() > c.Cap {
		v := c.List.Back()
		c.List.Remove(v)
		delete(c.Keys, v.Value.(Pair1).Key)
	}
}

type LRUCache2 struct {
	sync.RWMutex
	Shards map[string]*LRUCacheShard
}

type LRUCacheShard struct {
	LRU LRUCache1
	sync.RWMutex
}

func NewLRUCache2(capacity int) LRUCache2 {
	shards := make(map[string]*LRUCacheShard, 256)
	for i := 0; i < 256; i++ {
		shards[fmt.Sprintf("%02x", i)] = &LRUCacheShard{
			LRU: LRUCache1{
				Cap:  capacity,
				Keys: make(map[int]*list.Element),
				List: list.New(),
			},
		}
	}

	return LRUCache2{
		Shards: shards,
	}
}

func (c *LRUCache2) Get(key int) int {
	shard := c.GetShard(key)

	shard.RLock()
	defer shard.RUnlock()

	return shard.LRU.Get(key)
}

func (c *LRUCache2) Put(key, val int) {
	shard := c.GetShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.LRU.Put(key, val)
}

func (c *LRUCache2) GetShard(key int) (shard *LRUCacheShard) {
	hasher := sha1.New()
	hasher.Write([]byte{byte(key)})
	shardKey := fmt.Sprintf("%x", hasher.Sum(nil))[0:2]

	return c.Shards[shardKey]
}
