package main

func findComplement(num int) int {
	return revert2num(num2binaryOutput(num, ""), 0)
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

/*
1310. XOR Queries of a Subarray
You are given an array arr of positive integers. You are also given the array queries where queries[i] = [left[i], right[i]].
For each query i compute the XOR of elements from left[i] to right[i] (that is, arr[left[i]] XOR arr[left[i] + 1] XOR ... XOR arr[right[i]] ).
Return an array answer where answer[i] is the answer to the ith query.
*/
func xorQueries(arr []int, queries [][]int) []int {
	var output []int
	for _, index := range queries {
		currXOR := 0
		for i := index[0]; i <= index[1]; i++ {
			currXOR = currXOR ^ arr[i]
		}
		output = append(output, currXOR)
	}
	return output
}
