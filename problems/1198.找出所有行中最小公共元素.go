// 给你一个 m x n 的矩阵 mat，其中每一行的元素均符合 严格递增 。请返回 所有行中的 最小公共元素 。

// 如果矩阵中没有这样的公共元素，就请返回 -1。

// 输入：mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
// 输出：5

func smallestCommonElement(mat [][]int) int {
    n, m := len(mat), len(mat[0])
	pos := make([]int, n)
	curMax, cnt := 0, 0

	for {
		for i := 0; i < n; i++ {
			pos[i] = sort.Search(m, func(j int) bool {
				return mat[i][j] >= curMax
			})

			if pos[i] == m {
				return -1
			}
			if curMax == mat[i][pos[i]] {
				cnt++
				if cnt == n {
					return curMax
				}
			} else {
				cnt = 1
			}

			curMax = mat[i][pos[i]]
		}
	}
}