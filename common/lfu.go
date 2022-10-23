package common

type LFUCache struct {
	KeyToVal   map[int]int
	KeyToFreq  map[int]int
	FreqToKeys map[int][]int

	MinFreq int
	Cap     int
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		KeyToVal:   make(map[int]int),
		KeyToFreq:  make(map[int]int),
		FreqToKeys: make(map[int][]int),
		MinFreq:    0,
		Cap:        capacity,
	}
}

func (l *LFUCache) IncreaseFreq(key int) {
	freq := l.KeyToFreq[key]
	l.KeyToFreq[key] = freq + 1

	list := l.FreqToKeys[freq]
	for i, k := range list {
		if k == key {
			list = append(list[:i], list[i+1:]...)
		}
	}
	if len(list) == 0 {
		delete(l.FreqToKeys, freq)
		if l.MinFreq == freq {
			l.MinFreq++
		}
	} else {
		l.FreqToKeys[freq] = list
	}

	if _, ok := l.FreqToKeys[freq+1]; !ok {
		l.FreqToKeys[freq+1] = make([]int, 0)
	}
	l.FreqToKeys[freq+1] = append(l.FreqToKeys[freq+1], key)
}

func (l *LFUCache) RemoveMinFreqKey() {
	list := l.FreqToKeys[l.MinFreq]
	deleted := list[0]

	l.FreqToKeys[l.MinFreq] = list[1:]
	if len(l.FreqToKeys[l.MinFreq]) == 0 {
		delete(l.FreqToKeys, l.MinFreq)
	}

	delete(l.KeyToVal, deleted)
	delete(l.KeyToFreq, deleted)
}

func (l *LFUCache) Get(key int) int {
	if _, ok := l.KeyToVal[key]; !ok {
		return -1
	}

	l.IncreaseFreq(key)
	return l.KeyToVal[key]
}

func (l *LFUCache) Put(key, val int) {
	if _, ok := l.KeyToVal[key]; ok {
		l.KeyToVal[key] = val
		l.IncreaseFreq(key)
		return
	}

	if l.Cap <= len(l.KeyToVal) {
		l.RemoveMinFreqKey()
	}

	l.KeyToVal[key] = val
	l.KeyToFreq[key] = 1
	l.MinFreq = 1

	if _, ok := l.FreqToKeys[l.MinFreq]; !ok {
		l.FreqToKeys[l.MinFreq] = make([]int, 0)
	}
	l.FreqToKeys[l.MinFreq] = append(l.FreqToKeys[l.MinFreq], key)
}
