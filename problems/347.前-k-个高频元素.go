/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
type num struct {
	val  int
	freq int
}

type numHeap []num

func (h numHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h numHeap) Less(i, j int) bool  { return h[i].freq < h[j].freq }
func (h numHeap) Len() int            { return len(h) }
func (h *numHeap) Push(v interface{}) { *h = append(*h, v.(num)) }
func (h *numHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}

func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, n := range nums {
		seen[n]++
	}

	h := &numHeap{}
	for val, cnt := range seen {
		heap.Push(h, num{val, cnt})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, k)
	for i, num := range *h {
		ans[i] = num.val
	}
	return ans
}

// @lc code=end

