/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
type Trie struct {
	Children [26]*Trie
	Word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]
	}
	node.Word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	tree := &Trie{}
	for _, word := range words {
		tree.Insert(word)
	}

	result := make([]string, 0, len(words))

	for i, row := range board {
		for j, ch := range row {
			idx := ch - 'a'
			if tn := tree.Children[idx]; tn != nil {
				dfs(board, i, j, tree, &result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *Trie, result *[]string) {
	ch := board[i][j]
	node = node.Children[ch-'a']
	if node == nil {
		return
	}

	if node.Word != "" {
		*result = append(*result, node.Word)
		node.Word = ""
	}

	board[i][j] = '.'
	for _, dir := range dirs {
		nextx, nexty := i+dir.x, j+dir.y
		if 0 <= nextx && nextx < len(board) &&
			0 <= nexty && nexty < len(board[0]) &&
			board[nextx][nexty] != '.' {
			dfs(board, nextx, nexty, node, result)
		}
	}

	board[i][j] = ch
}

// @lc code=end

