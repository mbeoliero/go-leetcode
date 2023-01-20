/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	ans := 0
	root := &Trie{}
	for i := 1; i < len(nums); i++ {
		root.Add(nums[i-1])
		ans = max(ans, root.Check(nums[i]))
	}
	return ans
}

const highBit = 30

type Trie struct {
	Left, Right *Trie
}

func (t *Trie) Add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Left == nil {
				cur.Left = &Trie{}
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &Trie{}
			}
			cur = cur.Right
		}
	}
}

func (t *Trie) Check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.Right != nil {
				cur = cur.Right
				x = x*2 + 1
			} else {
				cur = cur.Left
				x = x * 2
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
				x = x*2 + 1
			} else {
				cur = cur.Right
				x = x * 2
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

