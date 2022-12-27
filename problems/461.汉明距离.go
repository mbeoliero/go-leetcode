/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	// Bitwise-XOR will give the number which have all the set bits for different corresponding btis
	xor := x ^ y
	// r := 0
	// Calculate the number of set bits
	// for xor != 0 {
	// 	xor -= xor & (-xor)
	// 	r++
	// }
	return bits.OnesCount(uint(xor))
}

// @lc code=end

