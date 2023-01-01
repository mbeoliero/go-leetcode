// 给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。
// 输入：intervals = [[0,30],[5,10],[15,20]]
// 输出：2
func minMeetingRooms(intervals [][]int) int {
	var nums []int
	for _, v := range intervals {
		nums = append(nums, v[0]*10+2)
		nums = append(nums, v[1]*10+1)
	}
	sort.Ints(nums)
	maxRoom := 0
	curNeedRoom := 0
	for _, v := range nums {
		if v%10 == 1 {
			curNeedRoom--
		} else {
			curNeedRoom++
		}
		if curNeedRoom > maxRoom {
			maxRoom = curNeedRoom
		}
	}
	return maxRoom
}

// solution 2
type hp struct {
	sort.IntSlice
}

func (h *hp) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() interface{} {
	v := h.IntSlice[:h.IntSlice.Len()-1]
	h.IntSlice = h.IntSlice[:h.IntSlice.Len()-1]
	return v
}

func minMeetingRooms(intervals [][]int) int {
	// 自定义排序+优先队列「最小堆」
	n := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &hp{}
	heap.Push(h, intervals[0][1])
	for i := 1; i < n; i++ {
		// 取出优先队列中最早结束的会议
		free := h.IntSlice[0]
		if free <= intervals[i][0] {
			// 证明前面会议已经结束，可以复用会议室
			heap.Pop(h)
		}
		// 将当前会议的结束时间加入优先队列
		heap.Push(h, intervals[i][1])
	}
	return h.IntSlice.Len()
}
