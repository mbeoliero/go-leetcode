package temp

import (
	"math"
	"sort"
)

func reverseWords(s string) string {
	r := []rune(s)
	reverseString(r, 0, len(s)-1)
	left, right := 0, 0
	for right < len(s) {
		if r[right] == 32 {
			reverseString(r, left, right-1)
			left = right
		}
		right++
	}
	reverseString(r, left, right-1)
	return string(r)
}

func reverseString(s []rune, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	left, right := 0, m-1
	upper, lower := 0, n-1
	res := make([]int, 0, m*n)

	for len(res) < m*n {
		if upper <= lower {
			for j := left; j <= right; j++ {
				res = append(res, matrix[upper][j])
			}
			upper++
		}

		if left <= right {
			for i := upper; i <= lower; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if upper <= lower {
			for j := right; j >= left; j-- {
				res = append(res, matrix[lower][j])
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

func removeDuplicateLetters(s string) string {
	stack := make([]rune, 0)
	count := [256]int{}
	for _, r := range s {
		count[r]++
	}

	for _, r := range s {
		for len(stack) > 0 && r > stack[0] && count[stack[0]] > 1 {
			count[stack[0]]--
			stack = stack[1:]
		}
		stack = append(stack, r)
	}
	return string(stack)
}

type RandomizedSet struct {
	Nums  []int
	Index map[int]int
}

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		Nums:  make([]int, 0),
// 		Index: make(map[int]int),
// 	}
// }

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Index[val]; ok {
		return false
	}
	this.Nums = append(this.Nums, val)
	this.Index[val] = len(this.Nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.Index[val]; ok {
		this.Nums = append(this.Nums[:v], this.Nums[v+1:]...)
		delete(this.Index, val)
		return true
	}
	return false
}

var memo [][]int
var keyMap map[byte][]int

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	keyMap = make(map[byte][]int)
	for i := 0; i < m; i++ {
		if _, ok := keyMap[ring[i]]; !ok {
			keyMap[ring[i]] = make([]int, 0)
		}
		keyMap[ring[i]] = append(keyMap[ring[i]], i)
	}

	return dp(ring, 0, key, 0)
}

var res int

func getNums(nums []int, index int) int {
	if index < 0 || index >= len(nums) {
		return 1
	}
	return nums[index]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	i := 0
	str := make([]byte, 0)
	for i < len(res) {
		str = append(str, byte('0'+res[i]))
		i++
	}
	return string(str)
}

func calculate(s string) int {
	stack := make([]int, 0)
	sign := '+'
	num := 0

	for i := 0; i < len(s); i++ {
		if isdigit(s[i]) {
			num = num*10 + int(s[i]-'0')
		}

		if !isdigit(s[i]) || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			}

			sign = rune(s[i])
			num = 0
		}
	}

	res := 0
	for i := len(stack) - 1; i >= 0; i-- {
		res += stack[i]
	}
	return res
}

func isdigit(c byte) bool {
	return c > '0' && c < '9'
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	bucket := make([]int, k)
	return backtrack(nums, bucket, target, 0)
}

func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, i, j int) bool {
	n := 9
	if j == n {
		return backtrack(board, i+1, 0)
	}
	if i == n {
		return true
	}

	if board[i][j] == '.' {
		return backtrack(board, i, j+1)
	}

	for ch := '1'; ch <= '9'; ch++ {
		if !isValid(board, i, j, byte(ch)) {
			continue
		}

		board[i][j] = byte(ch)
		if backtrack(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, ch byte) bool {
	n := 9
	for k := 0; k < n; k++ {
		if board[i][k] == ch || board[k][j] == ch {
			return false
		}
	}
	return true
}

// func heapify(nums []int, i, length int) {
// 	largest := i
// 	left, right := 2*i+1, 2*i+2

// 	if left < length && nums[largest] < nums[left] {
// 		largest = left
// 	}
// 	if right < length && nums[largest] < nums[right] {
// 		largest = right
// 	}

// 	if largest != i {
// 		swap(nums, i, largest)
// 		heapify(nums, largest, length)
// 	}
// }

func partition(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] > pivot {
			right--
		}
		nums[left] = nums[right]

		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left
}

// func findKthLargest(nums []int, k int) int {
// 	nlen := len(nums)
// 	p := partition(nums, 0, nlen-1)
// 	if nlen-p == k {
// 		return nums[p]
// 	}

// 	left, right := 0, nlen-1
// 	for nlen-p != k {
// 		if nlen-p > k {
// 			left = p + 1
// 		} else if nlen-p < k {
// 			right = p - 1
// 		}

// 		p = partition(nums, left, right)
// 	}
// 	return nums[p]
// }

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func heapify(nums []int, i int, nlen int) {
	small := i
	left, right := 2*i+1, 2*i+2
	if left < nlen && nums[small] > nums[left] {
		small = left
	}
	if right < nlen && nums[small] > nums[right] {
		small = right
	}

	if small != i {
		swap(nums, i, small)
		heapify(nums, small, nlen)
	}
}

func findKthLargest(nums []int, k int) int {
	nlen := len(nums)

	heap := make([]int, k)
	for i := 0; i < k; i++ {
		heap[0] = nums[i]
		heapify(heap, 0, k)
	}

	for i := k; i < nlen; i++ {
		if nums[i] > heap[0] {
			heap[0] = nums[i]
			heapify(heap, 0, k)
		}
	}

	return heap[0]
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrack func(nums []int, start int, path []int)
	backtrack = func(nums []int, start int, path []int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(nums, i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(nums, 0, []int{})
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	left := head
	res := true
	var traverse func(head *ListNode)

	traverse = func(head *ListNode) {
		if head == nil {
			return
		}

		traverse(head.Next)
		res = res && left.Val == head.Val
		left = left.Next
	}

	traverse(head)
	return res
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	st := []*TreeNode{}
	var cur *TreeNode = root

	for len(st) != 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}

		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if p == cur {
			break
		}

		cur = cur.Right
	}

	if len(st) == 0 {
		return nil
	}
	return st[len(st)-1]
}

func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, low, high int) bool
	dfs = func(root *TreeNode, low, high int) bool {
		if root == nil {
			return true
		}

		if root.Val >= high || root.Val <= low {
			return false
		}

		return dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, high)
	}

	return dfs(root, math.MinInt16, math.MaxInt16)
}

func isBalanced(root *TreeNode) bool {
	var res bool
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return 1
		}

		l := depth(root.Left)
		r := depth(root.Right)

		if l-r > 1 || r-l > 1 {
			res = false
		}
		return max(l, r) + 1
	}

	depth(root)
	return res
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := depth(root.Left)
	r := depth(root.Right)

	if l-r > 1 || r-l > 1 {

	}
	return max(l, r) + 1
}

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
