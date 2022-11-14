/*
 * @lc app=leetcode.cn id=864 lang=golang
 *
 * [864] 获取所有钥匙的最短路径
 */

// @lc code=start
func shortestPathAllKeys(grid []string) int {
	n, m := len(grid), len(grid[0])
	gridMatrix := make([][]int, n)

	start := []int{0, 0}
	count := 0

	for i := 0; i < n; i++ {
		gridMatrix[i] = make([]int, m)
		for j, r := range grid[i] {
			switch r {
			case '@':
				start = []int{i, j}
			case '#':
				gridMatrix[i][j] = -10
			case '.':
				gridMatrix[i][j] = 0
			default:
				if r >= 'a' && r <= 'f' {
					gridMatrix[i][j] = int(r-'a') + 1
					count++
				}
				if r >= 'A' && r <= 'F' {
					gridMatrix[i][j] = int(r-'A')*-1 - 1
				}
			}
		}
	}

	ops := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	canMove := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m && gridMatrix[i][j] != -10
	}

	type move struct {
		i, j  int
		keys  int
		steps int
	}
	getKey := func(m move) string {
		return fmt.Sprintf("%d_%d_%d", m.i, m.j, m.keys)
	}

	queue := make([]move, 0)
	queue = append(queue, move{start[0], start[1], 0, 0})
	visited := make(map[string]bool)

	// BFS
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.keys == (1<<count)-1 {
			return curr.steps
		}

		for _, op := range ops {
			nextI, nextJ := curr.i+op[0], curr.j+op[1]
			if canMove(nextI, nextJ) {
				next := move{nextI, nextJ, curr.keys, curr.steps + 1}

				if _, ok := visited[getKey(next)]; ok {
					continue
				}

				// lock
				if gridMatrix[nextI][nextJ] < 0 && gridMatrix[nextI][nextJ] > -10 {
					num := gridMatrix[nextI][nextJ] * -1
					if (1<<(num-1))&curr.keys == 0 {
						continue
					}
				}

				// key
				if gridMatrix[nextI][nextJ] > 0 {
					next.keys |= 1 << (gridMatrix[nextI][nextJ] - 1)
				}

				visited[getKey(next)] = true
				queue = append(queue, next)
			}
		}
	}

	return -1
}

// @lc code=end

