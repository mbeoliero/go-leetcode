package common

import "sort"

var a []int

type Hp struct{ sort.IntSlice }

func (h Hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *Hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *Hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
