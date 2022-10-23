/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

// @lc code=start
var ok = true
var color []bool
var visited []bool

func isBipartite(graph [][]int) bool {
	glen := len(graph)
	color = make([]bool, glen)
	visited = make([]bool, glen)

	for n := 0; n < glen; n++ {
		traverse(graph, n)
	}
	return ok
}

func traverse(graph [][]int, node int) {
	if !ok {
		return
	}

	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[node] {
			color[neighbor] = !color[node]
			traverse(graph, neighbor)
		} else {
			if color[neighbor] == color[node] {
				ok = false
				return
			}
		}
	}
}

// @lc code=end

