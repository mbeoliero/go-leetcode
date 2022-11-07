/*
 * @lc app=leetcode.cn id=939 lang=golang
 *
 * [939] 最小面积矩形
 */

// @lc code=start
func minAreaRect(points [][]int) int {
	hashX := map[int][]int{}
	for _, p := range points {
		hashX[p[0]] = append(hashX[p[0]], p[1])
	}
	allX := make([]int, 0, len(hashX))
	for k, v := range hashX {
		allX = append(allX, k)
		sort.Ints(v)
	}
	sort.Ints(allX)
	minArea := math.MaxInt32
	for i := 0; i < len(allX); i++ {
		for j := i + 1; j < len(allX); j++ {
			if temp := commonShortest(hashX[allX[i]], hashX[allX[j]]); temp != math.MaxInt32 {
				curr := temp * (allX[j] - allX[i])
				if curr < minArea {
					minArea = curr
				}
			}
		}
	}
	if minArea == math.MaxInt32 {
		return 0
	}
	return minArea
}

func commonShortest(s1, s2 []int) int {
	values := []int{}
	for i, j := 0, 0; i < len(s1) && j < len(s2); {
		if s1[i] == s2[j] {
			values = append(values, s1[i])
			i++
			j++
		} else if s1[i] > s2[j] {
			j++
		} else {
			i++
		}
	}
	shortest := math.MaxInt32
	for i := 0; i < len(values)-1; i++ {
		if values[i+1]-values[i] < shortest {
			shortest = values[i+1] - values[i]
		}
	}
	return shortest
}

// @lc code=end

func minAreaRect(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	left, right := 0, len(points)-1

	var res int
	for left < right {
		lh, rh := 0, 0

		tl := left + 1
		for points[tl][0] == points[left][0] {
			tl++
		}
		lh = points[tl-1][1] - points[left][1]

		tr := right - 1
		for points[tr][0] == points[right][0] {
			tr--
		}
		rh = points[right][1] - points[tr+1][1]

		width := points[right][0] - points[left][0]
		height := min(lh, rh)
		res = max(res, width*height)

		left, right = tl, tr
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
