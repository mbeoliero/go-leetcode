package common

type NumArray struct {
	Sum  []int
	Diff []int
}

func ConstructorSum(nums []int) NumArray {
	sums := make([]int, len(nums))
	for i, v := range nums {
		sums[i] = v
		if i > 0 {
			sums[i] += sums[i-1]
		}
	}
	return NumArray{
		Sum: sums,
	}
}

func (n *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return n.Sum[right]
	}
	return n.Sum[right] - n.Sum[left-1]
}

func ConstructorDiff(nums []int) NumArray {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return NumArray{
		Diff: diff,
	}
}

func (n *NumArray) DiffUpdate(left int, right int, update int) {
	n.Diff[left] += update
	if right+1 < len(n.Diff) {
		n.Diff[right+1] -= update
	}
}

func (n *NumArray) DiffResult() []int {
	nums := make([]int, len(n.Diff))
	nums[0] = n.Diff[0]
	for i := 1; i < len(n.Diff); i++ {
		nums[i] = nums[i-1] + n.Diff[i]
	}
	return nums
}
