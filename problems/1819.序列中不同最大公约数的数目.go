/*
 * @lc app=leetcode.cn id=1819 lang=golang
 *
 * [1819] 序列中不同最大公约数的数目
 */

// @lc code=start
func countDifferentSubsequenceGCDs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, x := range nums {
		if max < x {
			max = x
		}
	}

	g := make([]int, max+1)

	for _, x := range nums {
		for y := 1; y*y <= x; y++ {
			if x%y == 0 {
				if g[y] == 0 {
					g[y] = x
				} else {
					g[y] = gcd(g[y], x)
				}

				if y*y != x {
					z := x / y
					if g[z] == 0 {
						g[z] = x
					} else {
						g[z] = gcd(g[z], x)
					}
				}
			}
		}
	}

	ans := 0
	for i := 1; i <= max; i++ {
		if g[i] == i {
			ans++
		}
	}

	return ans
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

// @lc code=end

