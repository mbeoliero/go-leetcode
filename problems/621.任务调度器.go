/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	tsk := make([]int, 26)
	for _, t := range tasks {
		tsk[t-'A']++
	}
	max, count := len(tasks)/(n+1)+1, 0
	for _, v := range tsk {
		if v > max {
			max = v
			count = 1
		} else if v == max {
			count++
		}
	}
	if count == 0 {
		return len(tasks)
	}
	cpu := (max-1)*(n+1) + count
	// a trap here, cpu may less than len(tasks)
	// tasks=["A","A","A","B","B","C","D","E"], n=2
	// max=3, count=1, CPU=(3-1)*(2+1)+1=7, len(tasks)=8
	if cpu < len(tasks) {
		return len(tasks)
	}
	return cpu
}

// @lc code=end

