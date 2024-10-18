package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
2044. Count Number of Maximum Bitwise-OR Subsets
https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/description/
*/
func countMaxOrSubsets(nums []int) int {
	maxBitWise := 0
	for _, i := range nums {
		maxBitWise = maxBitWise | i
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		var arr []int
		arr = append(arr, nums[i])
		for j := i + 1; j < len(nums); j++ {
			for _, value := range arr {
				arr = append(arr, value|nums[j])
			}
		}
		for _, value := range arr {
			if value == maxBitWise {
				count++
			}
		}
	}
	return count
}

func num2binary(num int) string {
	if num == 0 {
		return "0"
	}
	if num == 1 {
		return "1"
	}
	if num%2 == 1 {
		return num2binary(num/2) + "1"
	}
	return num2binary(num/2) + "0"
}

func binaryToNum(binaryStr string) int {
	// Convert binary string to integer
	num, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error converting binary to number:", err)
		return 0
	}
	return int(num)
}

func num2binaryOutput(num int, output string) string {
	if num == 0 {
		return output
	}
	if num == 1 {
		return output + "1"
	}
	if num%2 == 1 {
		return num2binaryOutput(num/2, output) + "1"
	}
	return num2binaryOutput(num/2, output) + "0"
}

func revert2num(num string, output int) int {
	power := len(num)
	if power == 0 {
		return 0
	}
	if power == 1 {
		if num[0] == '0' {
			return output
		}
		return output + 1
	}
	if num[0] == '1' {
		return revert2num(num[1:], output+int(math.Pow(2, float64(power-1))))
	}
	return revert2num(num[1:], output)
}
