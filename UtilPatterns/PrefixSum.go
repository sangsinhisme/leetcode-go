package main

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return NumArray{sum: sum}
}

func (t *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return t.sum[right]
	}
	return t.sum[right] - t.sum[left-1]
}
