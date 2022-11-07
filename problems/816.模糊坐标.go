/*
 * @lc app=leetcode.cn id=816 lang=golang
 *
 * [816] 模糊坐标
 */

// @lc code=start
func ambiguousCoordinates(s string) []string {
	num := s[1 : len(s)-1]
	res := make([]string, 0)

	for l := 1; l < len(num); l++ {
		// lt := point(num, 0, l)
		lt := getPos(num[:l])
		if len(lt) == 0 {
			continue
		}
		// rt := point(num, l, len(num))
		rt := getPos(num[l:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return res
}

func point(num string, start, end int) []string {
	res := make([]string, 0)
	if num[start] != '0' || num[start:end] == "0" {
		res = append(res, num[start:end])
	}

	for i := start + 1; i < end; i++ {
		if i != start+1 && num[start] == '0' || num[end-1] == '0' {
			continue
		}
		res = append(res, num[start:i]+"."+num[i:end])
	}
	return res
}

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	for p := 1; p < len(s); p++ {
		if p != 1 && s[0] == '0' || s[len(s)-1] == '0' {
			continue
		}
		pos = append(pos, s[:p]+"."+s[p:])
	}
	return
}

// @lc code=end

func ambiguousCoordinates(s string) (res []string) {
	n := len(s) - 2
	s = s[1 : len(s)-1]
	for l := 1; l < n; l++ {
		lt := getPos(s[:l])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[l:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return
}