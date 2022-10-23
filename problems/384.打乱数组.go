/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		Nums: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.Nums
}

func (this *Solution) Shuffle() []int {
	arr := make([]int, len(this.Nums))
	copy(arr, this.Nums)

	for i := 0; i < len(this.Nums); i++ {
		ret := rand.Intn(len(this.Nums))
		arr[i], arr[ret] = arr[ret], arr[i]
	}
	return arr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

