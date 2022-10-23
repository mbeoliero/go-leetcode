package common

// UnionFindSet 并查集结构体
type UnionFindSet struct {
	People []int // 人员及其Header数组
	N      int   // 一共有多少人
}

func NewUnionFindSet(n int) *UnionFindSet {
	people := make([]int, n)
	// 让每一个人的父节点指向自己
	for i := range people {
		people[i] = i
	}

	return &UnionFindSet{
		People: people,
		N:      n,
	}
}

// Find 查找根节点
func (u *UnionFindSet) Find(x int) int {
	if u.People[x] == x {
		return x
	} else {
		// 如果他不是根节点，接着往上面找根节点，并把根节点赋给当前元素的父节点，构造二层的平铺树
		// 缩短查找距离
		u.People[x] = u.Find(u.People[x])
		return u.People[x]
	}
}

// Mix 合并两个节点到同一个联通域
func (u *UnionFindSet) Union(x, y int) {
	fx := u.Find(x)
	fy := u.Find(y)
	// 两个人不在一个联通域，把他们两个人的Header连起来
	if fx != fy {
		u.People[fx] = fy
	}
}
