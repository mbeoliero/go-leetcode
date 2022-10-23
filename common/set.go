package common

import (
	"sync"
)

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[interface{}]bool{},
	}
}

func (s *Set) Add(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Contains(item interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[interface{}]bool{}
}

func (s *Set) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.Len() == 0
}

func (s *Set) List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := []interface{}{}
	for k := range s.m {
		list = append(list, k)
	}
	return list
}

func (s *Set) Diff(t *Set) []interface{} {
	sDiff := []interface{}{}
	for sk := range s.m {
		if _, ok := t.m[sk]; !ok {
			sDiff = append(sDiff, sk)
		}
	}
	return sDiff
}
