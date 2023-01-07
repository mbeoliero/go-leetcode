// 给定一个 m x n 二进制数组表示的网格 grid ，一个岛屿由 四连通 （上、下、左、右四个方向）的 1 组成（代表陆地）。你可以认为网格的四周被海水包围。

// 如果两个岛屿的形状相同，或者通过旋转（顺时针旋转 90°，180°，270°）、翻转（左右翻转、上下翻转）后形状相同，那么就认为这两个岛屿是相同的。

// 返回 这个网格中形状 不同 的岛屿的数量 。

// 输入: grid = [[1,1,0,0,0],[1,0,0,0,0],[0,0,0,0,1],[0,0,0,1,1]]
// 1 1 0 0 0
// 1 0 0 0 0
// 0 0 0 0 1
// 0 0 0 1 1
// 输出: 1
// 解释: 这两个是相同的岛屿。因为我们通过 180° 旋转第一个岛屿，两个岛屿的形状相同。

import "sort"

func getKey(ars []*[]int) string {
	ss := make([]string, 8)
	for k := 0; k < 8; k++ {
		ar := *ars[k]
		sort.Ints(ar)
		y0, x0 := byte(ar[0]>>8), byte(ar[0]&0xff)
		bs := make([]byte, len(ar)<<1)
		for i, e := range ar {
			bs[i<<1], bs[i<<1+1] = byte(e>>8)+50-y0, byte(e&0xff)+50-x0
		}
		ss[k] = string(bs)
	}
	sort.Strings(ss)
	return ss[0]
}

func numDistinctIslands2(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var cal func(int, int, int, int, []*[]int)
	cal = func(i, j, i0, j0 int, ars []*[]int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		p2i := func(di, dj int) int {
			return ((di + 50) << 8) | (dj + 50)
		}
		di, dj := i-i0, j-j0
		*ars[0] = append(*ars[0], p2i(di, dj))
		*ars[1] = append(*ars[1], p2i(-di, dj))
		*ars[2] = append(*ars[2], p2i(di, -dj))
		*ars[3] = append(*ars[3], p2i(-di, -dj))
		*ars[4] = append(*ars[4], p2i(dj, di))
		*ars[5] = append(*ars[5], p2i(-dj, di))
		*ars[6] = append(*ars[6], p2i(dj, -di))
		*ars[7] = append(*ars[7], p2i(-dj, -di))
		cal(i-1, j, i0, j0, ars)
		cal(i, j-1, i0, j0, ars)
		cal(i+1, j, i0, j0, ars)
		cal(i, j+1, i0, j0, ars)
	}
	ms := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ars := make([]*[]int, 8)
			for k := 0; k < 8; k++ {
				ars[k] = &[]int{}
			}
			cal(i, j, i, j, ars)
			if 0 != len(*ars[0]) {
				ms[getKey(ars)] = true
			}
		}
	}
	return len(ms)
}

// 识别相似的岛屿：
// - 旋转90、180、270分别未(-y, x), (-x, -y), (y, -x)
// - 左右和上下翻转分别未(-x, y)和(x, -y)
// - 所有情况为(x, y), (x, -y), (-x, y), (-x, -y)以及起x、y互换
// - 对于所有的情况，每个点减去minx和miny，这样每种情况的点的集合是相同的
// - 为了hash到相同的string，可以将每个点转换为x * len(grid[0]) + y，保证每个转换为唯一的数字
// - 将转换后的slice排序，转换为string，选取所有情况中最大的string，即可识别相同形状岛屿

func numDistinctIslands2(grid [][]int) int {
	islandSet := map[string]struct{}{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			island := dfs(grid, i, j)
			if len(island) > 0 {
				str := hashfunc(island, len(grid[0]))
				islandSet[str] = struct{}{}
				// fmt.Println(island, str)
			}
		}
	}
	return len(islandSet)
}

var maxInt = 100000000

var direct = [][]int{
	{1, 1},
	{-1, -1},
	{1, -1},
	{-1, 1},
}

var dir = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxStr(a, b string) string {
	if a < b {
		return b
	} else {
		return a
	}
}

func hashfunc(iland [][]int, lift int) string {
	fmt.Println(iland)
	result := ""
	for i := 0; i < 4; i++ {
		minx1, miny1 := iland[0][0]*direct[i][0], iland[0][1]*direct[i][1]
		minx2, miny2 := maxInt, maxInt
		for _, point := range iland {
			minx1 = min(minx1, point[0]*direct[i][0])
			miny1 = min(miny1, point[1]*direct[i][1])
			minx2 = min(minx2, point[1]*direct[i][0])
			miny2 = min(miny2, point[0]*direct[i][1])
			// fmt.Println(minx2, miny2)
		}
		result1, result2 := []int{}, []int{}
		for _, point := range iland {
			result1 = append(result1, (point[0]*direct[i][0]-minx1)*lift+(point[1]*direct[i][1]-miny1))
			result2 = append(result2, (point[1]*direct[i][0]-minx2)*lift+(point[0]*direct[i][1]-miny2))
		}
		sort.Sort(sort.IntSlice(result1))
		sort.Sort(sort.IntSlice(result2))
		// fmt.Println(minx1, miny1, re sult1, minx2, miny2, result2)
		str1, str2 := "", ""
		for x := 0; x < len(result1); x++ {
			str1 += strconv.Itoa(result1[x])
			str2 += strconv.Itoa(result2[x])
		}
		if result == "" {
			result = maxStr(str1, str2)
		} else {
			result = maxStr(result, maxStr(str1, str2))
		}
	}
	return result
}

func siteLegal(x, y, mx, my int) bool {
	return x >= 0 && y >= 0 && x < mx && y < my
}

func dfs(grid [][]int, x, y int) [][]int {
	if grid[x][y] == 0 {
		return [][]int{}
	}
	grid[x][y] = 0
	result := [][]int{[]int{x, y}}
	for i := 0; i < 4; i++ {
		nx, ny := x+dir[i][0], y+dir[i][1]
		if siteLegal(nx, ny, len(grid), len(grid[0])) {
			result = append(result, dfs(grid, nx, ny)...)
		}
	}
	return result
}
