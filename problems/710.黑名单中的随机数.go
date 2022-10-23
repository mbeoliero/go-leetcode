/*
 * @lc app=leetcode.cn id=710 lang=golang
 *
 * [710] 黑名单中的随机数
 */

// @lc code=start
type Solution struct {
	whiteList map[int]int
	blackList map[int]struct{}
	n         int
	choice    bool //true for whiteList and false for blackList
}

func Constructor(n int, blacklist []int) Solution {
	s := Solution{
		whiteList: map[int]int{},
		blackList: map[int]struct{}{},
		n:         n,
		choice:    false,
	}
	// record blackList
	for _, v := range blacklist {
		s.blackList[v] = struct{}{}
	}
	if len(s.blackList) < n/5 { // here the bar could change for better practice
		return s
	}
	// build whiteList
	s.choice = true
	index := 0
	for i := 0; i < n; i++ {
		if _, ok := s.blackList[i]; !ok {
			s.whiteList[index] = i
			index++
		}
	}
	return s
}

func (this *Solution) Pick() int {
	if this.choice {
		return this.whiteList[rand.Intn(len(this.whiteList))]
	}

	for {
		ret := rand.Intn(this.n)
		if _, ok := this.blackList[ret]; !ok {
			return ret
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
// @lc code=end

