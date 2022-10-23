package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func findPalindrome(s string, idx int) string {
	p := idx
	for p >= 0 && p < len(s) {
		if s[p] != s[idx*2-p] {
			break
		}
		p--
	}
	return s[p+1 : idx*2-p]
}

func findPalindrome2(s string, idx1, idx2 int) string {
	if s[idx1] != s[idx2] {
		return ""
	}
	p := idx1
	for p >= 0 && p < len(s) {
		if s[p] != s[idx1*2+1-p] {
			break
		}
		p--
	}
	return s[p+1 : idx1*2+1-p]
}

func partition(nums []int, low, high int) int {
	target := nums[low]
	left, right := low+1, high
	for left < right {
		for nums[left] <= target {
			left++
		}
		for nums[right] > target {
			right--
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[low], nums[right] = nums[right], nums[low]
	return right
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt16
		for _, c := range coins {
			fmt.Println(i, c, dp[i])
			if i-c >= 0 {
				dp[i] = min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] == math.MaxInt16 {
		return -1
	}
	fmt.Println(dp)
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	tmp := math.MaxInt16
	var res int
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right-1 {
			sum := nums[left] + nums[right] + nums[i]
			if tmp != min(tmp, abs(target-sum)) {
				tmp = min(tmp, abs(target-sum))
				res = sum
			}
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
			if left == right-1 {
				break
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	ori, cnt := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		ori[s1[i]]++
	}

	for i := 0; i < len(s1); i++ {
		if _, ok := ori[s2[i]]; ok {
			cnt[s2[i]]++
		}
	}
	check := func() bool {
		for k, v := range ori {
			if cnt[k] != v {
				return false
			}
		}
		return true
	}

	left, right := 0, len(s1)
	for right <= len(s2) {
		fmt.Println(left, right)
		if check() {
			return true
		}

		if _, ok := ori[s2[right]]; ok {
			cnt[s2[right]]++
		}
		if _, ok := ori[s2[left]]; ok {
			cnt[s2[left]]--
		}
		left++
		right++
	}
	return false
}

func maxProfit(prices []int) int {
	var res int
	left, right := 0, 0
	for right < len(prices) {
		for right > 1 && right < len(prices) && prices[right] > prices[right-1] {
			right++
		}
		if right > 1 {
			res += prices[right-1] - prices[left]
		}
		left = right
		right++
	}
	return res
}

// func reverseString(s string) string {
// 	res := make([]rune, 0, len(s))
// 	func() {
// 		for _, v := range s {
// 			defer func(r rune) {
// 				res = append(res, r)
// 			}(v)
// 		}
// 	}()
// 	return string(res)
// }

// func reverseWords(s string) string {
// 	r := []rune(s)
// 	reverseString(r, 0, len(s)-1)
// 	fmt.Println(string(r))
// 	left, right := 0, 0

// 	r = append(r, 32)
// 	for right < len(r) {
// 		if r[right] == 32 {
// 			reverseString(r, left, right-1)
// 			left = right
// 		}
// 		right++
// 		for left < right {
// 			if r[left] == 32 {
// 				left++
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return string(r[:len(r)-1])
// }

// func reverseString(s []rune, left, right int) {
// 	for left < right {
// 		s[left], s[right] = s[right], s[left]
// 		left++
// 		right--
// 	}
// }

func reverseString(s []string, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// func spiralOrder(matrix [][]int) []int {
// 	m, n := len(matrix), len(matrix[0])
// 	i, j := 0, -1

// 	left, right := 0, m-1
// 	up, low := 0, n-1
// 	res := make([]int, 0, m*n)

// 	var x, y int
// 	for len(res) < m*n {
// 		x, y = getDirection(x, y)
// 		i += x
// 		j += y
// 		fmt.Println(x, y, i, j, left, right, up, low)

// 		for i >= left && i <= right && j >= up && j <= low {
// 			res = append(res, matrix[i][j])
// 			i += x
// 			j += y
// 		}

// 		i -= x
// 		j -= y

// 		switch {
// 		case x == 1:
// 			right -= x
// 		case x == -1:
// 			left -= x
// 		case y == 1:
// 			up += y
// 		case y == -1:
// 			low += y
// 		}
// 	}
// 	return res
// }

// func getDirection(x, y int) (int, int) {
// 	switch {
// 	case x == 0 && y == 1:
// 		return 1, 0
// 	case x == 1 && y == 0:
// 		return 0, -1
// 	case x == 0 && y == -1:
// 		return -1, 0
// 	default:
// 		return 0, 1
// 	}
// }

func removeDuplicateLetters(s string) string {
	cnt := make(map[rune][]int, 0)
	r := []rune(s)
	keys := make([]rune, 0)

	for i, v := range r {
		if _, ok := cnt[v]; !ok {
			cnt[v] = make([]int, 0)
			keys = append(keys, 0)
		}
		cnt[v] = append(cnt[v], i)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	res := make([]rune, len(s))
	p := len(s)
	for _, k := range keys {
		fd := p
		idx := cnt[k]
		for _, i := range idx {
			if i > p {
				fd = i
				break
			}
		}
		if fd == p {
			res[idx[len(idx)-1]] = k
		} else {
			res[fd] = k
		}
	}

	n := make([]rune, 0)
	for _, v := range res {
		if v != 0 {
			n = append(n, v)
		}
	}
	return string(n)
}

func findRepeatedDnaSequences(s string) []string {
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'G':
			nums[i] = 1
		case 'C':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}
	fmt.Println(nums)

	L, R := 10, 4
	RL := int(math.Pow(float64(R), float64(L-1)))
	fmt.Println(RL)

	hash := make(map[int]bool)
	ap := make(map[int]bool)
	res := make([]string, 0)
	wHash := 0

	left, right := 0, 0
	for right < len(nums) {
		wHash = R*wHash + nums[right]
		right++
		fmt.Println(wHash, right)

		if right-left == L {
			fmt.Println(hash, ap, wHash)
			if hash[wHash] && !ap[wHash] {
				res = append(res, s[left:right])
				ap[wHash] = true
			} else {
				hash[wHash] = true
			}
			wHash -= nums[left] * RL
			left++
		}
	}
	return res
}

func gen() <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			ch <- n
			n++
			fmt.Println(n)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	// nums := []int{1, 2, 5}
	// i := 0
	// j := len(nums) - 1
	// for i < j {
	// 	nums[i], nums[j] = nums[j], nums[i]
	// 	i++
	// 	j--
	// }
	// fmt.Println(nums)
	// println(findPalindrome2("cbbd", 1, 2))
	//
	// p := make([]int, len(nums))
	// copy(p, nums)
	// res := make([]int, 0, len(nums))
	// res = append(res, nums...)
	// nums[2] = 1
	// fmt.Println(p, nums, res)
	// fmt.Println(threeSumClosest([]int{0, 1, 2}, 3))
	// fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	// c := common.ConstructorDiff([]int{0, 0, 0, 0, 0})
	// c.DiffUpdate(1, 3, 2)
	// c.DiffUpdate(2, 4, 3)
	// c.DiffUpdate(0, 2, -2)
	// fmt.Println(c.DiffResult())
	// s := "  hello   world  "
	// strList := strings.Fields(s)
	// reverseString(strList, 0, len(strList)-1)
	// fmt.Println(strings.Join(strList, " "))
	// s := "bcabc"
	// r := []rune(s)
	// r[0] = 0
	// last := len(r) - 1
	// for last > 0 {
	// 	fmt.Println(r, last)
	// 	r = r[:last]
	// 	fmt.Println(len(r))
	// 	fmt.Println(r)
	// 	break
	// }
	// sort.Slice([]int{1, 2}, func(i, j int) bool {})
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 3; i++ {
	// 	println(rand.Intn(100))
	// }
	// nums := []int{1, 4, 4}
	// fmt.Println(function(nums, 1))
	// fmt.Println(math.Pow(2, 10))
	// nums := []int{2, 2}
	// target := 2
	// s := make([]*int, 0)
	// i, j := 1, 2
	// s = append(s, &i, nil, &j)
	// fmt.Println(s, len(s))
	// fmt.Println(searchRange(nums, target))
	// fmt.Println(rightBound(nums, target, function))
	// fmt.Println(strings.Split("1,2,3,null,null,4,5", ","))
	// fmt.Println(HeapSort([]int{3, 2, 1, 5, 6, 4}))
	// a := "sret"
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 19}}))
	// fmt.Println(pancakeSort([]int{1, 2, 3}))
	// for n := range gen() {
	// 	fmt.Println("main: ", n)
	// 	if n == 5 {
	// 		break
	// 	}
	// }
	// time.Sleep(time.Minute)
	// fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	// t := 1 << 3
	// s := 16 >> 2 & 1
	fmt.Printf("%T", '1')
}

func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		left, right := i-1, i+1
		if height[left] > height[i] && height[right] > height[i] {
			fmt.Println("i: res: ", i, res)
			j := left
			for ; j >= 0 && height[j] >= height[j+1]; j-- {
			}
			left = j + 1

			j = right
			for ; j < len(height)-1 && height[j] >= height[j-1]; j++ {
			}
			right = j - 1

			fmt.Println(left, right)
			temp := 0
			for j = left + 1; j < right; j++ {
				temp += height[j]
			}
			fmt.Println(temp)

			res += min(height[left], height[right])*(right-left-1) - temp
		}
	}
	return res
}
