/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 */

// @lc code=start
type NumArray struct {
	Sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	for i, v := range nums {
		sums[i] = v
		if i > 0 {
			sums[i] += sums[i-1]
		}
	}
	return NumArray{
		Sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.Sums[right]
	}
	return this.Sums[right] - this.Sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

