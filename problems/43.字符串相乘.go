/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
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
	for i < len(res) && res[i] == 0 {
		i++
	}
	for i < len(res) {
		str = append(str, byte('0'+res[i]))
		i++
	}
	if len(str) == 0 {
		return "0"
	}
	return string(str)
}

// @lc code=end

