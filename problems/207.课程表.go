/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// func canFinish(numCourses int, prerequisites [][]int) bool {
// 	var queue []int

// 	indegree := make(map[int][]int)
// 	adjacent := make(map[int][]int)
// 	count := 0

// 	for _, prerequisite := range prerequisites {
// 		src := prerequisite[1]
// 		dst := prerequisite[0]

// 		indegree[dst] = append(indegree[dst], src)
// 		adjacent[src] = append(adjacent[src], dst)
// 	}

// 	for i := 0; i < numCourses; i++ {
// 		if len(indegree[i]) == 0 {
// 			enqueue(&queue, i)
// 		}
// 	}

// 	for len(queue) > 0 {
// 		dequeuedEle := dequeue(&queue)

// 		for _, vertex := range adjacent[dequeuedEle] {
// 			tmp := indegree[vertex]
// 			remove(&tmp, dequeuedEle)
// 			indegree[vertex] = tmp

// 			if len(indegree[vertex]) == 0 {
// 				enqueue(&queue, vertex)
// 			}
// 		}

// 		count += 1
// 	}

// 	if count == numCourses {
// 		return true
// 	}

// 	return false
// }

// func remove(lst *[]int, removedEle int) {
// 	if lst == nil {
// 		panic("nil pointer")
// 	}

// 	if len(*lst) == 0 {
// 		panic("empty list")
// 	}

// 	for idx, num := range *lst {
// 		if num == removedEle {
// 			*lst = append((*lst)[:idx], (*lst)[idx+1:]...)
// 		}
// 	}
// }

// func enqueue(queue *[]int, newEle int) {
// 	if queue == nil {
// 		panic("nil pointer")
// 	}

// 	*queue = append(*queue, newEle)
// }

// func dequeue(queue *[]int) int {
// 	if queue == nil {
// 		panic("nil pointer")
// 	}

// 	if len(*queue) == 0 {
// 		panic("empty queue")
// 	}

// 	dequeuedEle := (*queue)[0]

// 	*queue = (*queue)[1:]

// 	return dequeuedEle
// }

// @lc code=end

