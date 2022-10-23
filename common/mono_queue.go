package common

type MonotonicQueue struct { //创建结构体
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue { //初始化函数
	return &MonotonicQueue{
		queue: make([]int, 0),
	}
}

func (m *MonotonicQueue) Push(k int) {
	//将单调队列中比k小的全都弹出
	if len(m.queue) > 0 && k > m.queue[0] { //队列中所有的元素都比k小
		m.queue = []int{}
	}
	for len(m.queue) > 0 && m.queue[len(m.queue)-1] < k { // 队列中所可能存在元素比k小
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, k)
}

func (m *MonotonicQueue) Pop(t int) {
	if m.queue[0] == t { //判断滑动窗口最左边元素t是不是单调队列最大值，若是弹出，否则什么都不用做
		m.queue = m.queue[1:]
	}
}
