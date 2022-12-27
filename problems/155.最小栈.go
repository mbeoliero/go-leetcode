/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{
		Stack: make([]int, 0),
		Min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.Stack) == 1 {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, min(val, this.Min[len(this.Min)-1]))
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return -1
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		return -1
	}
	return this.Min[len(this.Min)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

