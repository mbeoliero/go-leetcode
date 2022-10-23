/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
const inf = math.MaxInt64 / 2

var memo [][]int

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([][]int, n)
	for i := range prices {
		prices[i] = make([]int, n)
		for j := range prices[i] {
			prices[i][j] = inf
		}
	}

	memo = make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+2)
		for j := range memo[i] {
			memo[i][j] = -888
		}
	}

	for _, i := range flights {
		prices[i[0]][i[1]] = i[2]
	}

	if dp(prices, src, dst, k+1) == inf {
		return -1
	}
	return dp(prices, src, dst, k+1)
}

func dp(prices [][]int, src, dst, k int) int {
	if src == dst {
		return 0
	}
	if k == 0 {
		return inf
	}

	if memo[src][k] != -888 {
		return memo[src][k]
	}

	res := prices[src][dst]
	for i, p := range prices[src] {
		if p != inf && dp(prices, i, dst, k-1) != inf {
			res = min(res, p+dp(prices, i, dst, k-1))
		}
	}

	memo[src][k] = res
	return res
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

// @lc code=end

const MAX_INT = int(^uint(0) >> 1)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	graph := make([][][]int, n)
	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], flight[1:])
	}
	h := &NodeHeap{}
	heap.Init(h)
	heap.Push(h, &Node{src, 0, 0})
	for h.Len() != 0 {
		currNode := heap.Pop(h).(*Node)
		currId := currNode.id
		if currId == dst {
			return currNode.dist
		}
		// 2. update the neighbors and heap.
		if currNode.stops <= K { // !!!
			for _, edge := range graph[currId] {
				heap.Push(h, &Node{edge[0], currNode.dist + edge[1], currNode.stops + 1})
			}
		}
	}
	return -1
}

type Node struct{ id, dist, stops int }
type NodeHeap []*Node

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h NodeHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h NodeHeap) Len() int            { return len(h) }
func (h *NodeHeap) Push(v interface{}) { *h = append(*h, v.(*Node)) }
func (h *NodeHeap) Pop() interface{} {
	ret := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return ret
}