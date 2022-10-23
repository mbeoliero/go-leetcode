/*
 * @lc app=leetcode.cn id=528 lang=golang
 *
 * [528] 按权重随机选择
 */

// @lc code=start
type Solution struct {
	Vals []int
	Sum  int
}

func Constructor(w []int) Solution {
	sum := 0
	vals := make([]int, 0)
	for _, weight := range w {
		sum += weight
		vals = append(vals, sum)
	}
	return Solution{
		Vals: vals,
		Sum:  sum,
	}
}

func (this *Solution) PickIndex() int {
	r := rand.Intn(this.Sum)
	return this.Find(r)
}

func (this *Solution) Find(r int) int {
	left, right := 0, len(this.Vals)
	for left <= right {
		mid := left + (right-left)/2
		if this.Vals[mid] == r {
			return mid + 1
		} else if this.Vals[mid] < r {
			left = mid + 1
		} else if this.Vals[mid] > r {
			right = mid - 1
		}
	}
	return left
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
// @lc code=end

