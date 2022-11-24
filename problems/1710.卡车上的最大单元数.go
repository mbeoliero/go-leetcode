/*
 * @lc app=leetcode.cn id=1710 lang=golang
 *
 * [1710] 卡车上的最大单元数
 */

// @lc code=start
func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	sum := 0
	res := 0
	for _, box := range boxTypes {
		if sum == truckSize {
			break
		}
		num := box[0]
		if sum+num > truckSize {
			res += (truckSize - sum) * box[1]
			sum = truckSize
		} else {
			res += num * box[1]
			sum += num
		}
	}

	return res
}

// @lc code=end

