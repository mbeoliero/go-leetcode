/*
 * @lc app=leetcode.cn id=1860 lang=golang
 *
 * [1860] 增长的内存泄露
 */

// @lc code=start
func memLeak(memory1 int, memory2 int) []int {
	memory := []int{memory1, memory2}
	getIndex := func() int {
		index := 0
		if memory[1] > memory[0] {
			index = 1
		}
		return index
	}

	i := 1
	for {
		index := getIndex()
		if i > memory[index] {
			return []int{i, memory[0], memory[1]}
		}

		memory[index] -= i
		i++
	}

	return []int{}
}

// @lc code=end

