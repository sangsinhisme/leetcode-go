package main

import "strconv"

/*
670. Maximum Swap
https://leetcode.com/problems/maximum-swap/description/
*/
func maximumSwap(num int) int {
	s := strconv.Itoa(num)
	nums := make([]int, len(s))
	for i := range nums {
		nums[i] = int(s[i])
	}
	swap1 := -1
	swap2 := -1
	maxSwapPossible := -1
	for i := len(s) - 1; i > -1; i-- {
		if maxSwapPossible == -1 || nums[i] > nums[maxSwapPossible] {
			maxSwapPossible = i
		} else if nums[i] < nums[maxSwapPossible] {
			swap1 = i
			swap2 = maxSwapPossible
		}
	}
	if swap1 != -1 && swap2 != -1 {
		nums[swap2], nums[swap1] = nums[swap1], nums[swap2]
		output := 0
		for _, i := range nums {
			output = output*10 + (i - 48)
		}
		return output
	} else {
		return num
	}
}
