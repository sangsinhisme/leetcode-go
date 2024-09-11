package main

import "math"

func findComplement(num int) int {
	return revert2num(num2binary(num, ""), 0)
}

func num2binary(num int, output string) string {
	if num == 0 {
		return output
	}
	if num == 1 {
		return output + "1"
	}
	if num%2 == 1 {
		return num2binary(num/2, output) + "1"
	}
	return num2binary(num/2, output) + "0"
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

/*
2220. Minimum Bit Flips to Convert Number
A bit flip of a number x is choosing a bit in the binary representation of x and flipping it from either 0 to 1 or 1 to 0.

For example, for x = 7, the binary representation is 111 and we may choose any bit (including any leading zeros not shown)
and flip it. We can flip the first bit from the right to get 110, flip the second bit from the right to get 101, flip the
fifth bit from the right (a leading zero) to get 10111, etc.
Given two integers start and goal, return the minimum number of bit flips to convert start to goal.
*/
func minBitFlips(start int, goal int) int {
	var helper func(num int, diff int) int
	helper = func(num int, diff int) int {
		if num == 0 {
			return diff
		}
		if num == 1 {
			return diff + 1
		}
		if num%2 == 1 {
			return helper(num/2, diff+1)
		}
		return helper(num/2, diff)
	}
	return helper(start^goal, 0)
}
