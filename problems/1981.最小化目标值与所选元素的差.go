/*
 * @lc app=leetcode.cn id=1981 lang=golang
 *
 * [1981] 最小化目标值与所选元素的差
 */

// @lc code=start
func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	sum := 0

	magic := make([][]int, m)
	for i := 0; i < m; i++ {
		sort.Ints(mat[i])
		sum += mat[i][0]

		for j := 1; j < n; j++ {
			if mat[i][j] != mat[i][j-1] {
				magic[i] = append(magic[i], mat[i][j]-mat[i][0])
			}
		}
	}

	minDiff := abs(target - sum)
	if sum >= target {
		return minDiff
	}

	hash := make(map[int]bool, 871)
	hash[sum] = true

	for _, tires := range magic {
		more := make([]int, 0)

		for _, tryValue := range tires {
			for basic := range hash {
				if basic+tryValue > 870 {
					continue
				}

				if hash[basic+tryValue] == false {
					more = append(more, basic+tryValue)
				}
				if basic+tryValue == target {
					return 0
				}
			}
		}

		for _, mv := range more {
			hash[mv] = true
		}
	}

	for k := range hash {
		if curDiff := abs(target - k); curDiff < minDiff {
			minDiff = curDiff
		}
	}

	return minDiff
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

// @lc code=end

