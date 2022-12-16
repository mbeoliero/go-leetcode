/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, w := range word {
		ch := w - 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
}

func wordBreak(s string, wordDict []string) bool {
	root := &Trie{}
	for _, word := range wordDict {
		root.Insert(word)
	}
	failed := make([]bool, 301)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if failed[i] {
			return false
		}

		if i == len(s) {
			return true
		}

		node := root
		for j := i; j < len(s); j++ {
			ch := s[j] - 'a'
			if node.Children[ch] != nil {
				node = node.Children[ch]
				if node.IsEnd && dfs(j+1) {
					return true
				}
			} else {
				break
			}
		}
		failed[i] = true
		return false
	}

	return dfs(0)
}

// @lc code=end

