/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})

	res := 0
	curEnd, nextEnd := 0, 0
	for i := 0; i < len(clips) && clips[i][0] <= curEnd; {
		for i < len(clips) && clips[i][0] <= curEnd {
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		res++
		curEnd = nextEnd

		if curEnd >= time {
			return res
		}
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

