/*
 * @lc app=leetcode.cn id=990 lang=golang
 *
 * [990] 等式方程的可满足性
 */

// @lc code=start
func equationsPossible(equations []string) bool {
	u := new(Union)
	u.Init()

	notEqual := make([][]int, 0)
	for _, equal := range equations {
		a, b := int(equal[0]-'a'), int(equal[3]-'a')
		if equal[1] == '=' {
			u.Union(a, b)
		} else {
			notEqual = append(notEqual, []int{a, b})
		}
	}

	for _, not := range notEqual {
		a, b := not[0], not[1]
		pa, pb := u.Find(a), u.Find(b)
		if pa == pb {
			return false
		}
	}
	return true
}

type Union struct {
	Parent []int
}

func (u *Union) Init() {
	u.Parent = make([]int, 26)
	for i := range u.Parent {
		u.Parent[i] = i
	}
}

func (u *Union) Find(i int) int {
	if i == u.Parent[i] {
		return i
	} else {
		u.Parent[i] = u.Find(u.Parent[i])
		return u.Parent[i]
	}
}

func (u *Union) Union(a, b int) {
	pa := u.Find(a)
	pb := u.Find(b)
	if pa != pb {
		u.Parent[pb] = pa
	}
}

// @lc code=end

