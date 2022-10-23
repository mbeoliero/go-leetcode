/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */

// @lc code=start
var res [][]int
var end int

func allPathsSourceTarget(graph [][]int) [][]int {
	res = make([][]int, 0)
	end = len(graph) - 1

	traverse(graph, 0, []int{})
	return res
}

func traverse(graph [][]int, i int, path []int) {
	path = append(path, i)
	if i == end {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		// path = path[:len(path)-1]
		return
	}

	for _, g := range graph[i] {
		traverse(graph, g, path)
	}

	path = path[:len(path)-1]
}

// @lc code=end

