/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 */

// @lc code=start
import "math/rand"

type RandomizedSet struct {
	Nums  []int
	Index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Nums:  make([]int, 0),
		Index: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Index[val]; ok {
		return false
	}
	this.Nums = append(this.Nums, val)
	this.Index[val] = len(this.Nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.Index[val]; !ok {
		return false
	}
	last := this.Nums[len(this.Nums)-1]
	if last != val {
		idx := this.Index[val]
		this.Nums[idx] = last
		this.Index[last] = idx
	}
	this.Nums = this.Nums[:len(this.Nums)-1]
	delete(this.Index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.Nums))
	return this.Nums[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end

